package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vllvll/keepa/gen"
	"github.com/vllvll/keepa/internal/types"
)

func (k *KeepaServer) CreateBinary(ctx context.Context, in *pb.CreateBinaryRequest) (*pb.BinaryResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	binary := types.Binary{
		Meta:    in.GetMeta(),
		Content: in.GetBinary(),
	}

	binaryId, err := k.binaryRepository.Create(binary, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	binaryResponse := pb.BinaryResponse{
		Meta:      binary.Meta,
		Id:        binaryId,
		Binary:    binary.Content,
		UpdatedAt: timestamppb.Now(),
	}

	return &binaryResponse, nil
}
