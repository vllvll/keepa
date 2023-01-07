package handlers

import (
	"context"
	"github.com/vllvll/keepa/internal/types"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k *KeepaServer) UpdateBankCard(ctx context.Context, in *pb.UpdateBankCardRequest) (*pb.BankCardResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	bankCard := types.BankCard{
		ID:     in.GetId(),
		Number: in.GetNumber(),
		Holder: in.GetHolder(),
		CVV:    in.GetCvv(),
		Meta:   in.GetMeta(),
	}

	err := k.bankCardRepository.Update(bankCard)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	bankCard, err = k.bankCardRepository.Get(in.GetId())
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
