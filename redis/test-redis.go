package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)

	var err error
	err = client.Set("key", "xushaozhang2012", 0).Err()
	if err != nil {
		fmt.Println("set key, err:", err)
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		fmt.Println("get key, err:", err)
		panic(err)
	}
	fmt.Println("get key:", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
