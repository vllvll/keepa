package handlers

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/vllvll/keepa/gen"
)

func (k *KeepaServer) Login(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	login := in.GetLogin()
	password := in.GetPassword()

	if login == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty login")
	}

	if password == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty password")
	}

	user, err := k.userRepository.GetUserHashByLogin(login)
	if err != nil || !k.cryptService.IsEqual(password, user.Hash) {

		return nil, status.Error(codes.InvalidArgument, "Incorrect password")
	}

	token, err := k.cryptService.GenerateRand()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	err = k.tokenRepository.CreateToken(token, user.ID)
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
