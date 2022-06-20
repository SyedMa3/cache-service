package db

import (
	"log"
	"strconv"
	"time"

	pb "github.com/SyedMa3/cache-service/z_generated"

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
		log.Fatalln(err)
	}

	return key, nil
}

func (r *RedisDB) Get(key string) ([]byte, error) {
	var val []byte
	err := Rcache.Get(r.client.Context(), "mateen:"+key, &val)
	if err != nil {
		log.Fatalln(err)
	}

	return val, nil
}

func (r *RedisDB) SetUser(in *pb.SetUserRequest) (string, error) {
	key := "mateen:" + in.GetName() + strconv.FormatInt(in.GetRollNum(), 10)
	err := Rcache.Set(&cache.Item{
		Ctx: r.client.Context(),
		Key: key,
		Value: pb.SetUserRequest{
			Name:     in.GetName(),
			Class:    in.GetClass(),
			RollNum:  in.GetRollNum(),
			Metadata: in.GetMetadata(),
		},
		TTL: time.Second * 300,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return key, nil
}

func (r *RedisDB) GetUser(in *pb.GetUserRequest) (*pb.UserResponse, error) {

	key := "mateen:" + in.GetName() + strconv.FormatInt(in.GetRollNum(), 10)
	// log.Println(key)
	var val *pb.UserResponse
	err := Rcache.Get(r.client.Context(), key, &val)

	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}

	Rcache = cache.New(&cache.Options{
		Redis: rdb,
	})

	return &RedisDB{client: rdb}, nil
}
