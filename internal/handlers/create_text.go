package handlers

import (
	"context"
	"github.com/vllvll/keepa/internal/types"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	text, err = k.textRepository.Get(textId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	bankCardResponse := pb.TextResponse{
		Meta:      text.Meta,
		Id:        text.ID,
		Text:      text.Content,
		UpdatedAt: timestamppb.New(text.UpdatedAt),
	}

	return &bankCardResponse, nil
}
