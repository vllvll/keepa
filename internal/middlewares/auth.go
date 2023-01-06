package middlewares

import (
	"context"
	"github.com/vllvll/keepa/internal/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Auth(userRepository repositories.UserInterface, tokenRepository repositories.TokenInterface) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var token string
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			values := md.Get("keepa-auth")
			if len(values) > 0 {
				token = values[0]
			}
		}

		userID, err := tokenRepository.GetUserIDByToken(token)
		if err != nil {
			return handler(ctx, req)
		}

		user, err := userRepository.GetUserByID(userID)
		if err != nil {
			return handler(ctx, req)
		}

		contextWithUser := context.WithValue(ctx, "user", user)

		return handler(contextWithUser, req)
	}
}
