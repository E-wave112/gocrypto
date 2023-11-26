package pkg

import (
	"context"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	users, resultErr := redisClient.LRange(ctx, "subscribers", 0, -1).Result()
	if resultErr != nil {
		// panic(resultErr)
		log.Fatalf("an error occurred while retrieving the users %s\n", resultErr)
	}

	return users
}

func SetValueInRedis(email string) (string, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// setErr := redisClient.Set(ctx, "foo", "bar", 0).Err()
	// if setErr != nil {
	// 	fmt.Printf("an error occurred while setting your redis key: %s\n", setErr)
	// }

	subscribedUsers := retrieveValuesFromRedisStore()
	log.Printf("subscribed users count %v", len(subscribedUsers))
	if len(subscribedUsers) == 0 {
		setErr := redisClient.LPush(ctx, "subscribers", email).Err()
		if setErr != nil {
			// panic(setErr)
			log.Fatalf("an error occurred while inserting the initial value: %s\n", setErr)
		}
		response := fmt.Sprintf("%s successfully subscribed", email)
		return response, true
	}
	// check if email already exists in the redis list (insert if false, otherwise return the list as is)
	filteredSubscribedUsers := slices.DeleteFunc[[]string](subscribedUsers, func(item string) bool {
		return item != email

	})

	log.Printf("filtered subs %v", filteredSubscribedUsers)

	if len(filteredSubscribedUsers) == 0 {
		setErr := redisClient.LPush(ctx, "subscribers", email).Err()
		if setErr != nil {
			log.Fatalf("an error occurred while inserting the value: %s\n", setErr)
		}
		response := fmt.Sprintf("%s successfully subscribed", email)
		return response, true
	}

	return "email already exists", false

}

func UnsetValueInRedis(email string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	subscribedUsers := retrieveValuesFromRedisStore()
	log.Printf("subscribed users count %v", len(subscribedUsers))

	if len(subscribedUsers) == 0 {
		return "this email address is not subscribed to this alert", false
	}

	// check if email already exists in the redis list (return error if false, otherwise offset from the list)
	filteredSubscribedUsers := slices.DeleteFunc[[]string](subscribedUsers, func(item string) bool {
		return item != email

	})
	log.Printf("filtered subs %v", filteredSubscribedUsers)

	if len(filteredSubscribedUsers) == 0 {
		return "this email address is not subscribed to this alert", false
	}

	unSetErr := redisClient.LRem(ctx, "subscribers", 0, email).Err()
	if unSetErr != nil {
		log.Fatalf("an error occurred while unsetting the value: %s\n", unSetErr)
	}
	response := fmt.Sprintf("%s successfully unsubscribed", email)
	return response, true

}
