package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var RCache *redis.Client

func init() {
	// Initialize the redis connection to a redis instance running on your local machine
	password := os.Getenv("REDIS_PASSWORD")
	RCache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	err := ping(RCache)
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to Redis")
	}
}

func ping(client *redis.Client) error {
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}
