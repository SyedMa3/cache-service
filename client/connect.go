package main

import (
	"log"

	pb "github.com/SyedMa3/cache-service/z_generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddr = "localhost:9000"

func main() {

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect at %v", serverAddr)
	}
	defer conn.Close()

	client := pb.NewRpcServiceClient(conn)

	TestRawConn(client)
	TestUserConn(client)
}
