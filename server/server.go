package main

import (
	"context"
	"log"
	"net"

	"github.com/SyedMa3/cache-service/db"
	pb "github.com/SyedMa3/cache-service/proto"

	"google.golang.org/grpc"
)

var DB *db.RedisDB

type rpcServer struct {
	pb.UnimplementedRpcServiceServer
}

func (s *rpcServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.Response, error) {
	val, err := DB.Get(in.GetKey())
	if err != nil {

	}
	return &pb.Response{Value: val}, nil
}

func (s *rpcServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.Response, error) {
	return &pb.Response{Value: []byte("not implemented yet. Mateen will implement me")}, nil
}

func newServer() *rpcServer {
	s := &rpcServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen at port 9000")
	}

	grpcServer := grpc.NewServer()

	DB, err = db.CreateDB()
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterRpcServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
