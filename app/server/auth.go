package server

import (
	"context"
	"fmt"

	pb "github.com/noahharshbarger/game-node/server/auth"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Printf("Registering user: %s\n", req.Username)
	return &pb.RegisterResponse{
		Message: "User registered successfully",
		Success: true,
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("Logging in user: %s\n", req.Username)
	return &pb.LoginResponse{
		Token: "some-jwt-token",
		Success: true,
	}, nil
}