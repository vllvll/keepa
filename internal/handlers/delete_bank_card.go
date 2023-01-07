package handlers

import (
	"context"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (k *KeepaServer) DeleteBankCard(ctx context.Context, in *pb.DeleteBankCardRequest) (*emptypb.Empty, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	bankCardId := in.GetId()

	err := k.bankCardRepository.Delete(bankCardId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	return new(emptypb.Empty), nil
}
