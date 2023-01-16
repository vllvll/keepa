// Package handlers содержит обработчики проекта
package handlers

import (
	"context"

	pb "github.com/vllvll/keepa/gen"
	"github.com/vllvll/keepa/internal/repositories"
	"github.com/vllvll/keepa/internal/services"
	"github.com/vllvll/keepa/internal/types"
)

// KeepaServer поддерживает все необходимые методы сервера.
type KeepaServer struct {
	pb.UnimplementedKeepaServer

	userRepository     repositories.UserInterface
	tokenRepository    repositories.TokenInterface
	bankCardRepository repositories.BankCardInterface
	textRepository     repositories.TextInterface
	binaryRepository   repositories.BinaryInterface
	cryptService       services.CryptInterface
}

// NewHandler Получение хендлера
func NewHandler(
	userRepository repositories.UserInterface,
	tokenRepository repositories.TokenInterface,
	bankCardRepository repositories.BankCardInterface,
	textRepository repositories.TextInterface,
	binaryRepository repositories.BinaryInterface,
	cryptService services.CryptInterface,
) *KeepaServer {
	return &KeepaServer{
		userRepository:     userRepository,
		tokenRepository:    tokenRepository,
		bankCardRepository: bankCardRepository,
		textRepository:     textRepository,
		binaryRepository:   binaryRepository,
		cryptService:       cryptService,
	}
}

// Handlers Список методов для хендлеров (сервер)
type Handlers interface {
	Login(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error)
	Register(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error)
}

func GetUser(ctx context.Context) *types.User {
	userValue := ctx.Value("user")
	if userValue == nil {
		return nil
	}

	user, _ := userValue.(types.User)

	return &user
}
