package handlers

import (
	"context"

	"github.com/SyedMa3/cache-service/server/logic"
	pb "github.com/SyedMa3/cache-service/z_generated"
)

type RpcServer struct {
	pb.UnimplementedRpcServiceServer
}

func (s *RpcServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.Response, error) {
	val, err := logic.RDB.Get(in.GetKey())
	if err != nil {
		return &pb.Response{Status: err.Error()}, err
	}
	return &pb.Response{Value: val, Status: "Success"}, nil
}

func (s *RpcServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.Response, error) {
	_, err := logic.RDB.Set(in.GetKey(), in.GetValue())
	if err != nil {
		return &pb.Response{Status: err.Error()}, err
	}
	return &pb.Response{Value: in.GetValue(), Status: "Success"}, nil
}
