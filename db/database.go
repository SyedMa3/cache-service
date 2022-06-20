package db

import (
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type DB interface {
	Set(key string, value []byte) (string, error)
	Get(key string) (string, error)
}

var Rcache *cache.Cache

type RedisDB struct {
	client *redis.Client
}

func (r *RedisDB) Set(key string, value []byte) (string, error) {
	err := Rcache.Set(&cache.Item{
		Ctx:   r.client.Context(),
		Key:   "mateen:" + key,
		Value: value,
		TTL:   time.Second * 300,
	})

	if err != nil {
		log.Fatal(err)
	}

	return key, nil
}

func (r *RedisDB) Get(key string) ([]byte, error) {
	var val []byte
	err := Rcache.Get(r.client.Context(), key, &val)
	if err != nil {
		log.Fatal(err)
	}

	return val, nil
}

func CreateDB() (*RedisDB, error) {
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

	return &RedisDB{client: rdb}, nil
}
