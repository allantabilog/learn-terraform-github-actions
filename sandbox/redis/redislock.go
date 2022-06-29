package main

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {

	// Connect to redis
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})

	defer client.Close()

	// Create a new lock client
	locker := redislock.New(client)

	ctx := context.Background()

	// Try to obtain the lock
	lock, err := locker.Obtain(ctx, "name", 100*time.Millisecond, nil)

	if err == redislock.ErrNotObtained {
		log.Fatal("Could not obtain lock")
	} else if err != nil {
		log.Fatalln(err)
	}

	// defer release of the lock
	defer lock.Release(ctx)

	fmt.Println("Obtained the lock")

	// sleep and check the remaining TTL
	time.Sleep(50 * time.Millisecond)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl > 0 {
		fmt.Println("TTL still > 0, still possessing lock")
	}

	// extend my lock
	if err := lock.Refresh(ctx, 100*time.Millisecond, nil); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(50 * time.Millisecond)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl == 0 {
		fmt.Println("The lock has expired")
	}

}
