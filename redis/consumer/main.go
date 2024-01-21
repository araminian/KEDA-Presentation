package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Replace these values with your Redis server information
	redisAddr := getEnv("REDIS_ADDR", "192.168.194.170:6379")
	redisPassword := getEnv("REDIS_PASSWORD", "V21D68b9qN")
	redisDB := 0

	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Ping the Redis server to check the connection
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	// Key of the Redis list
	listKey := "usernames"

	fmt.Println("Waiting for data from Redis list...")

	for {
		time.Sleep(10 * time.Second)
		remainingUsers, err := getRedisListSize(client, listKey)
		if err != nil {
			log.Fatalf("Error getting size of Redis list: %v", err)
		}
		fmt.Println("Remaining users in Redis list:", remainingUsers)
		if remainingUsers == 0 {
			fmt.Println("No users remaining in Redis list. Waiting...")
			continue
		}

		// Pop data from the Redis list
		data, err := popFromRedisList(client, listKey)
		if err != nil {
			log.Fatalf("Error popping data from Redis list: %v", err)
		}

		fmt.Printf("User popped from Redis list: %s\n", data)
		fmt.Println("Processing data...")

	}
}

// addToRedisList adds data to a Redis list
func addToRedisList(client *redis.Client, key string, data interface{}) error {

	// LPush adds elements to the head (left) of the list
	_, err := client.LPush(key, data).Result()
	return err
}

// pop from the Redis list
func popFromRedisList(client *redis.Client, key string) (string, error) {

	// RPop removes the last element from the list and returns it
	return client.RPop(key).Result()
}

// Get size of Redis list
func getRedisListSize(client *redis.Client, key string) (int64, error) {

	// LLen returns the length of the list stored at key
	return client.LLen(key).Result()
}
