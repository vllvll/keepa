package handlers

import (
	"context"
	"github.com/vllvll/keepa/internal/types"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k *KeepaServer) UpdateBinary(ctx context.Context, in *pb.UpdateBinaryRequest) (*pb.BinaryResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	binary := types.Binary{
		ID:      in.GetId(),
		Content: in.GetBinary(),
		Meta:    in.GetMeta(),
	}

	err := k.binaryRepository.Update(binary)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	binary, err = k.binaryRepository.Get(in.GetId())
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
