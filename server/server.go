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
		return &pb.Response{Status: err.Error()}, err
	}
	return &pb.Response{Value: val, Status: "Success"}, nil
}

func (s *rpcServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.Response, error) {
	_, err := DB.Set(in.GetKey(), in.GetValue())
	if err != nil {
		return &pb.Response{Status: err.Error()}, err
	}
	return &pb.Response{Value: in.GetValue(), Status: "Success"}, nil
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
