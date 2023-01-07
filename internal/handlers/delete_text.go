package handlers

import (
	"context"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (k *KeepaServer) DeleteText(ctx context.Context, in *pb.DeleteTextRequest) (*emptypb.Empty, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	textID := in.GetId()

	err := k.textRepository.Delete(textID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	return new(emptypb.Empty), nil
}
