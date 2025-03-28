package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/skyfora/go-redis/internal"
)

func SetLogger(logger internal.Logging) {
	internal.Logger = logger
}

type Client interface {
	Set(ctx context.Context, key string, value interface{}) error
	SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	GetList(ctx context.Context, key string) ([]string, error)
	SetList(ctx context.Context, key string, values []string) error
	AddToList(ctx context.Context, key string, value string) error
	RemoveFromList(ctx context.Context, key string, value string) error
	Delete(ctx context.Context, key string) error
	DeletePattern(ctx context.Context, prefix string) error
	FlushAll(ctx context.Context) error
}

type client struct {
	client *redis.Client
	ttl    time.Duration
}

func New(URL string, ttlSecond int) Client {
	c := redis.NewClient(&redis.Options{
		Addr: URL,
	})

	_, err := c.Ping(context.Background()).Result()
	if err != nil {
		internal.Logger.Printf(context.Background(), "failed to connect to redis, error %s", err.Error())
	}

	return &client{
		client: c,
		ttl:    time.Duration(ttlSecond),
	}
}

func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	err := c.client.Set(ctx, key, value, c.ttl*time.Second)
	if err.Err() != nil {
		internal.Logger.Printf(context.Background(), "failed to set key %s, error %s", key, err.Err().Error())
	}
	return err.Err()
}

func (c *client) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *client) GetList(ctx context.Context, key string) ([]string, error) {
	vals, err := c.client.ZRangeByLex(ctx, key, &redis.ZRangeBy{Min: "-", Max: "+"}).Result()
	if err != nil {
		return nil, err
	}
	return vals, nil
}

func (c *client) SetList(ctx context.Context, key string, values []string) error {
	pipe := c.client.Pipeline()
	for _, v := range values {
		pipe.ZAdd(ctx, key, redis.Z{Score: 0, Member: v}) // Same score for lexicographic ordering
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (c *client) AddToList(ctx context.Context, key string, value string) error {
	return c.client.ZAdd(ctx, key, redis.Z{Score: 0, Member: value}).Err()
}

func (c *client) RemoveFromList(ctx context.Context, key string, value string) error {
	return c.client.ZRem(ctx, key, value).Err()
}

func (c *client) Delete(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key)
	if err.Err() != nil {
		internal.Logger.Printf(context.Background(), "failed to delete key %s, error %s", key, err.Err().Error())
	}
	return err.Err()
}

func (c *client) DeletePattern(ctx context.Context, prefix string) error {
	iter := c.client.Scan(
		ctx,
		0,
		prefix+"*",
		0,
	).Iterator()

	for iter.Next(ctx) {
		err := c.client.Del(
			ctx,
			iter.Val(),
		).Err()
		if err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}

func (c *client) FlushAll(ctx context.Context) error {
	return c.client.FlushAll(ctx).Err()
}
