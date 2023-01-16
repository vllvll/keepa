package handlers

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/vllvll/keepa/gen"
)

func (k *KeepaServer) Register(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	login := in.GetLogin()
	password := in.GetPassword()

	if login == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty login")
	}

	if password == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty password")
	}

	if k.userRepository.IsExists(login) {
		return nil, status.Error(codes.InvalidArgument, "Login is exists")
	}

	passwordHash := k.cryptService.Hash(password)

	userID, err := k.userRepository.CreateUser(login, passwordHash)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	token, err := k.cryptService.GenerateRand()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	err = k.tokenRepository.CreateToken(token, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	header := metadata.Pairs("keepa-auth", token)
	err = grpc.SendHeader(ctx, header)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	response := new(pb.AuthResponse)
	response.Token = token

	return response, nil
}
