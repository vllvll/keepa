// Package handlers содержит обработчики проекта
package handlers

import (
	"context"
	"github.com/vllvll/keepa/internal/repositories"
	"github.com/vllvll/keepa/internal/services"
	pb "github.com/vllvll/keepa/proto"
)

// KeepaServer поддерживает все необходимые методы сервера.
type KeepaServer struct {
	pb.UnimplementedKeepaServer

	userRepository  repositories.UserInterface
	tokenRepository repositories.TokenInterface
	cryptService    services.CryptInterface
}

// NewHandler Получение хендлера
func NewHandler(
	userRepository repositories.UserInterface,
	tokenRepository repositories.TokenInterface,
	cryptService services.CryptInterface,
) *KeepaServer {
	return &KeepaServer{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		cryptService:    cryptService,
	}
}

// Handlers Список методов для хендлеров (сервер)
type Handlers interface {
	Login(_ context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error)
	Register(_ context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error)
}
