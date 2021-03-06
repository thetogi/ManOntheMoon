package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
)

var RCache *redis.Client

func init() {
	// Initialize the redis connection to a redis instance running on your local machine

	address := os.Getenv("REDIS_ADDRESS")
	port := os.Getenv("REDIS_PORT")
	database, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	password := os.Getenv("REDIS_PASSWORD")
	RCache = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password, // no password set
		DB:       database, // use default DB
	})

	err := ping(RCache)
	if err != nil {
		panic("REDIS: " + err.Error())
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
