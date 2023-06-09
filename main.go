package creatingasimpleredisclientingo

import (
	"context"
	"fmt"
	"github.com/AydenRosinkranz/simple-redis-client/cache"
	"log"
)

const (
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDB       int    = 0
)

func main() {
	fmt.Println("hello world..")
	ctx := context.Background()
	c := cache.New(RedisAddr, RedisPassword, RedisDB)
	if err := c.Ping(ctx); err != nil {
		log.Panic("failed to connect to Redis")
	}

	log.Print("connected to Redis ..")

	if err := c.Set(ctx, "user:name", "AydenRosinkranz", 0); err != nil {
		log.Println("Error:could not store a value in Redis")
	}

	log.Println("value stored in Redis")

	res, err := c.Get(ctx, "user:name")
	if err != nil {
		log.Print("Error: could not get a value from Redis")
	}

	log.Print("Result: %s\n", res)
}
