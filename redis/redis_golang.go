package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var RedisClient *redis.Client

func ConnectRedis() *redis.Client {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDBStr := os.Getenv("REDIS_DB")

	// Convert redisDBStr to int
	redisDB, err := strconv.Atoi(redisDBStr)
	if err != nil {
		log.Fatal("Error converting REDIS_DB to int")
	}

	// Create Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Check connection to Redis
	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)
	return RedisClient
}
