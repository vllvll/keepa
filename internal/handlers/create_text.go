package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vllvll/keepa/gen"
	"github.com/vllvll/keepa/internal/types"
)

func (k *KeepaServer) CreateText(ctx context.Context, in *pb.CreateTextRequest) (*pb.TextResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	text := types.Text{
		Meta:    in.GetMeta(),
		Content: in.GetText(),
	}

	textId, err := k.textRepository.Create(text, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	bankCardResponse := pb.TextResponse{
		Meta:      text.Meta,
		Id:        textId,
		Text:      text.Content,
		UpdatedAt: timestamppb.Now(),
	}

	return &bankCardResponse, nil
}
