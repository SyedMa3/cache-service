package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SyedMa3/cache-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddr = "localhost:9000"

func testConn(client pb.RpcServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, &pb.GetRequest{Key: "test"})
	if err != nil {
		log.Fatalf("client.Get Failed!")
	}
	log.Printf("client.Get response: %s", resp)

	resp, err = client.Set(ctx, &pb.SetRequest{Key: "test", Value: []byte("testvalue")})
	if err != nil {
		log.Fatalf("client.Set Failed!")
	}
	log.Printf("client.Set resonse: %s", resp)
}

func main() {

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect")
	}
	defer conn.Close()

	client := pb.NewRpcServiceClient(conn)

	testConn(client)
}
