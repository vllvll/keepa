package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/vllvll/keepa/gen"
)

func (k *KeepaServer) DeleteBankCard(ctx context.Context, in *pb.DeleteBankCardRequest) (*emptypb.Empty, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	bankCardId := in.GetId()

	err := k.bankCardRepository.Delete(bankCardId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	return new(emptypb.Empty), nil
}
