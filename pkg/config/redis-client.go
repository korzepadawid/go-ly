package config

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	db, _ := strconv.Atoi(GetEnv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_ADDR"),
		Password: GetEnv("REDIS_PASSWORD"),
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalln("Can't connect with redis.")
	}

	log.Println("Connected with redis.")

	Redis = client
}
