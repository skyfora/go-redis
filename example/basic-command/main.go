package main

import (
	"context"
	"log"

	"github.com/skyfora/go-redis"
)

func main() {
	r := redis.New("localhost:6379", 60)
	log.Println("connected to redis")

	log.Println("setting key1")
	err := r.Set(context.Background(), "key1", "value")
	if err != nil {
		panic(err)
	}

	log.Println("setting key2")
	err = r.Set(context.Background(), "key2", "value")
	if err != nil {
		panic(err)
	}

	valueStr, err := r.Get(context.Background(), "key1")
	if err != nil {
		panic(err)
	}
	log.Printf("key1: %s\n", valueStr)

	log.Println("deleting key*")
	err = r.DeletePattern(context.Background(), "key")
	if err != nil {
		panic(err)
	}

	_, err = r.Get(context.Background(), "key1")
	if err != redis.Nil {
		panic(err)
	}
	log.Printf("key1 empty")

	log.Println("program end")
}
