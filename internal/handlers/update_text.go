package handlers

import (
	"context"
	"github.com/vllvll/keepa/internal/types"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k *KeepaServer) UpdateText(ctx context.Context, in *pb.UpdateTextRequest) (*pb.TextResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	text := types.Text{
		ID:      in.GetId(),
		Content: in.GetText(),
		Meta:    in.GetMeta(),
	}

	err := k.textRepository.Update(text)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	text, err = k.textRepository.Get(in.GetId())
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
