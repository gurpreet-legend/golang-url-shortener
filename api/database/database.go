package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	pong, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Print("Unable to connect to database")
	} else {
		fmt.Print("Connection established!", pong)
	}
	return rdb
}
