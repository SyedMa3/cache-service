package main

import (
	"log"
	"net"

	"github.com/SyedMa3/cache-service/server/handlers"
	"github.com/SyedMa3/cache-service/server/logic"
	pb "github.com/SyedMa3/cache-service/z_generated"
	"google.golang.org/grpc"
)

func newServer() *handlers.RpcServer {

	s := &handlers.RpcServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen at port 9000")
	}

	gRpcServer := grpc.NewServer()

	err = logic.CreateDB()
	if err != nil {
		log.Fatalln(err)
	}

	pb.RegisterRpcServiceServer(gRpcServer, newServer())
	gRpcServer.Serve(lis)
}
