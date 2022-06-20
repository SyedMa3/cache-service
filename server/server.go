package main

import (
	"context"
	"log"
	"net"

	pb "github.com/SyedMa3/cache-service/proto"

	"google.golang.org/grpc"
)

type rpcServer struct {
	pb.UnimplementedRpcServiceServer
}

func (s *rpcServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{Value: []byte("Mateen will implement me")}, nil
}

func (s *rpcServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.Response, error) {
	return &pb.Response{Value: []byte("not implemented yet. Mateen will implement me")}, nil
}

func newServer() *rpcServer {
	s := &rpcServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", "9000")
	if err != nil {
		log.Fatalf("Failed to listen")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterRpcServiceServer(grpcServer, newServer())

	grpcServer.Serve(lis)

}
