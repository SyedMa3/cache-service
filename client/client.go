package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SyedMa3/cache-service/z_generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddr = "localhost:9000"

func testConn(client pb.RpcServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Set(ctx, &pb.SetRequest{Key: "test", Value: []byte("testvalue")})
	if err != nil {
		log.Fatalf("client.Set Failed!")
	}
	log.Println(resp)

	resp, err = client.Get(ctx, &pb.GetRequest{Key: "test"})
	if err != nil {
		log.Fatalf("client.Get Failed!")
	}
	log.Println(resp)

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

func main() {

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect at %v", serverAddr)
	}
	defer conn.Close()

	client := pb.NewRpcServiceClient(conn)

	testConn(client)
}
