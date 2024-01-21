package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
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

	numberOfUsers := 10

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

	// Data to be added to the Redis list
	usernames := generateUsernames(numberOfUsers)

	// Add each username to the Redis list
	for _, username := range usernames {

		// Add data to the Redis list
		err = addToRedisList(client, listKey, username)
		if err != nil {
			log.Fatalf("Error adding data to Redis list: %v", err)
		}

		fmt.Printf("Data added to Redis list: %s\n", username)
	}
}

// addToRedisList adds data to a Redis list
func addToRedisList(client *redis.Client, key string, data interface{}) error {

	// LPush adds elements to the head (left) of the list
	_, err := client.LPush(key, data).Result()
	return err
}

// Generate a list of random usernames
func generateUsernames(numbers int) []string {
	var usernames []string
	for i := 0; i < numbers; i++ {
		// Generate a random number between 0 and 1000
		randNum := rand.Intn(10000)
		// Generate hash of current time
		hash := md5.New()
		hash.Write([]byte(time.Now().String()))
		// get 10 characters from the hash
		username := fmt.Sprintf("%x", hash.Sum(nil))[:10]

		// Append username to the slice
		usernames = append(usernames, fmt.Sprintf("username-%s-%d", username, randNum))
	}
	return usernames
}
