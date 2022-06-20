package logic

import (
	"log"
	"strconv"
	"time"

	pb "github.com/SyedMa3/cache-service/z_generated"

	"github.com/go-redis/cache/v8"
)

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
