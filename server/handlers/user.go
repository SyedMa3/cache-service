package handlers

import (
	"context"

	"github.com/SyedMa3/cache-service/server/logic"
	pb "github.com/SyedMa3/cache-service/z_generated"
)

func (s *RpcServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	val, err := logic.RDB.GetUser(in)
	if err != nil {
		return &pb.UserResponse{Status: err.Error()}, err
	}
	return val, nil
}

func (s *RpcServer) SetUser(ctx context.Context, in *pb.SetUserRequest) (*pb.UserResponse, error) {
	val, err := logic.RDB.SetUser(in)
	if err != nil {
		return &pb.UserResponse{Status: err.Error()}, err
	}
	return &pb.UserResponse{Key: val, Status: "Success"}, nil
}
