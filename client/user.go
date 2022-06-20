package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SyedMa3/cache-service/z_generated"
)

func TestUserConn(client pb.RpcServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userResp, err := client.SetUser(ctx, &pb.SetUserRequest{
		Name:     "Test Name",
		Class:    "Testing Class",
		RollNum:  111,
		Metadata: []byte("s"),
	})
	if err != nil {
		log.Fatalln("client.SetUser Failed!")
	}
	log.Println(userResp)

	userResp, err = client.GetUser(ctx, &pb.GetUserRequest{
		Name:    "Test Name",
		RollNum: 111,
	})
	if err != nil {
		log.Fatalln("client.GetUser Failed")
	}
	log.Println(userResp)
}
