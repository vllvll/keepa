package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vllvll/keepa/gen"
)

func (k *KeepaServer) GetText(ctx context.Context, in *pb.GetTextRequest) (*pb.TextResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	id := in.GetId()

	text, err := k.textRepository.Get(id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	textResponse := pb.TextResponse{
		Meta:      text.Meta,
		Id:        text.ID,
		Text:      text.Content,
		UpdatedAt: timestamppb.New(text.UpdatedAt),
	}

	return &textResponse, nil
}
