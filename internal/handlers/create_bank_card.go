package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vllvll/keepa/gen"
	"github.com/vllvll/keepa/internal/types"
)

func (k *KeepaServer) CreateBankCard(ctx context.Context, in *pb.CreateBankCardRequest) (*pb.BankCardResponse, error) {
	user := GetUser(ctx)
	if user == nil {
		return nil, status.Error(codes.PermissionDenied, "Unauthorized")
	}

	bankCard := types.BankCard{
		Number: in.GetNumber(),
		Holder: in.GetHolder(),
		CVV:    in.GetCvv(),
		Meta:   in.GetMeta(),
	}

	id, err := k.bankCardRepository.Create(bankCard, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	bankCardResponse := pb.BankCardResponse{
		Meta:      bankCard.Meta,
		Id:        id,
		Number:    bankCard.Number,
		Holder:    bankCard.Holder,
		Cvv:       bankCard.CVV,
		UpdatedAt: timestamppb.Now(),
	}

	return &bankCardResponse, nil
}
