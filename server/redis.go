package server

import (
	"context"
	"log"
	"slices"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func client() *redis.Client {
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
		log.Fatalf("Invalid enviroment variable!")
	}
	opt, err := redis.ParseURL(value)
	if err != nil {
		log.Fatalf("an error occurred while setting your redis instance: %s\n", err)

	}
	client := redis.NewClient(opt)

	return client
}

var redisClient = client()

func retrieveValuesFromRedisStore() []string {

	users, resultErr := redisClient.LRange(ctx, "subscribers", 0, -1).Result()
	if resultErr != nil {
		// panic(resultErr)
		log.Fatalf("an error occurred while retrieving the users %s\n", resultErr)
	}
	return users
}

func setValueInRedis(email string) (string, bool) {

	// setErr := redisClient.Set(ctx, "foo", "bar", 0).Err()
	// if setErr != nil {
	// 	fmt.Printf("an error occurred while setting your redis key: %s\n", setErr)
	// }

	subscribedUsers := retrieveValuesFromRedisStore()
	if len(subscribedUsers) == 0 {
		setErr := redisClient.LPush(ctx, "subscribers", email).Err()
		if setErr != nil {
			// panic(setErr)
			log.Fatalf("an error occurred while inserting the initial value: %s\n", setErr)
		}
		return "success", true
	}
	// check if email already exists in the redis list (insert if false, otherwise return the list as is)
	filteredSubscribedUsers := slices.DeleteFunc[[]string](subscribedUsers, func(item string) bool {
		return item != email

	})

	log.Printf("filtered subs %v", filteredSubscribedUsers)

	if len(filteredSubscribedUsers) == 0 {
		setErr := redisClient.LPush(ctx, "subscribers", email).Err()
		if setErr != nil {
			// panic(setErr)
			log.Fatalf("an error occurred while inserting the value: %s\n", setErr)
		}
	}

	return "email already exists", false

}
