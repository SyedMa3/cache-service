package logic

import (
	"log"
	"time"

	pb "github.com/SyedMa3/cache-service/z_generated"
	"github.com/go-redis/cache/v8"
)

type DB interface {
	Set(key string, value []byte) (string, error)
	Get(key string) (string, error)
	SetUser(in *pb.SetRequest) (string, error)
	GetUser(in *pb.GetUserRequest) (*pb.UserResponse, error)
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
