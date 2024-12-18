# Go Redis Client for Skyfora

A reusable redis client for Skyfora's Go project.

## Installation

Follow the instructions from [Go Official Website](https://go.dev/doc/faq#git_https)

Configure the ~/.sshconfig to use SSH on https url, by adding the following:

```bash
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```

Or use password authentication by modifying the ~/.netrc file:

```
machine github.com
login <username>
password <password>
```

Then, you will also need to modify the GOPRIVATE environment variable:

```bash
(For Linux/MacOS)
export GOPRIVATE=github.com/skyfora/*
or
go env -w GOPRIVATE='github.com/skyfora/*'
```

Instalation on the project:

```bash
go get github.com/skyfora/go-redis
```

## Usage

Details in the [example](example)

```go
import (
    "github.com/skyfora/go-redis"
)

func main() {
	r := redis.New("localhost:6379", 60)

	err := r.Set(context.Background(), "key1", "value")
	if err != nil {
		panic(err)
	}
}
```
