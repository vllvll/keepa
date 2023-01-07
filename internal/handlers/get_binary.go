package handlers

import (
	"context"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k *KeepaServer) GetBinary(ctx context.Context, in *pb.GetBinaryRequest) (*pb.BinaryResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	id := in.GetId()

	binary, err := k.binaryRepository.Get(id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	binaryResponse := pb.BinaryResponse{
		Meta:      binary.Meta,
		Id:        binary.ID,
		Binary:    binary.Content,
		UpdatedAt: timestamppb.New(binary.UpdatedAt),
	}

	return &binaryResponse, nil
}
