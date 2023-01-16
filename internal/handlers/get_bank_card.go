package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vllvll/keepa/gen"
)

func (k *KeepaServer) GetBankCard(ctx context.Context, in *pb.GetBankCardRequest) (*pb.BankCardResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	id := in.GetId()

	bankCard, err := k.bankCardRepository.Get(id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	bankCardResponse := pb.BankCardResponse{
		Meta:      bankCard.Meta,
		Id:        bankCard.ID,
		Number:    bankCard.Number,
		Holder:    bankCard.Holder,
		Cvv:       bankCard.CVV,
		UpdatedAt: timestamppb.New(bankCard.UpdatedAt),
	}

	return &bankCardResponse, nil
}
