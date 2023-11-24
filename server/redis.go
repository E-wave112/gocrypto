package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func client() string {
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get("REDIS_URL").(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid enviroment variable")
	}
	opt, err := redis.ParseURL(value)
	if err != nil {
		fmt.Printf("an error occurred while setting your redis instance: %s\n", err)

	}
	client := redis.NewClient(opt)

	setErr := client.Set(ctx, "foo", "bar", 0).Err()
	if setErr != nil {
		fmt.Printf("an error occurred while setting your redis key: %s\n", setErr)
	}
	val, valErr := client.Get(ctx, "foo").Result()
	if valErr != nil {
		fmt.Printf("an error occurred while retrieving the redis value: %s\n", valErr)
	}
	print(val)
	return val
}
