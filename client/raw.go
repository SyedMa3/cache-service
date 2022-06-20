package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SyedMa3/cache-service/z_generated"
)

func TestRawConn(client pb.RpcServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Set(ctx, &pb.SetRequest{Key: "test", Value: []byte("testvalue")})
	if err != nil {
		log.Fatalf("client.Set Failed!, %v", err)
	}
	log.Println(resp)

	resp, err = client.Get(ctx, &pb.GetRequest{Key: "test"})
	if err != nil {
		log.Fatalf("client.Get Failed!")
	}
	log.Println(resp)
}
