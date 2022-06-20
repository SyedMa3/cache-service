package logic

import (
	"log"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var Rcache *cache.Cache

type RedisDB struct {
	client *redis.Client
}

var RDB *RedisDB

func CreateDB() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatalln(err)
	}

	Rcache = cache.New(&cache.Options{
		Redis: rdb,
	})

	RDB = &RedisDB{client: rdb}

	return nil
}
