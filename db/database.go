package db

import (
	"log"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var Rcache *cache.Cache

func createDB() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	Rcache = cache.New(&cache.Options{
		Redis: rdb,
	})

	return rdb, nil
}
