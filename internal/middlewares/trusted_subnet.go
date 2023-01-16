package middlewares

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	conf "github.com/vllvll/keepa/internal/config"
)

func TrustSubnet() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		config, err := conf.CreateServerConfig()
		if err != nil {
			log.Fatalf("Error with config: %v", err)
		}

		var ip string

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			values := md.Get("ip")
			if len(values) > 0 {
				ip = values[0]
			}
		}

		if config.TrustedSubnet != "" {
			if ip == "" {
				return nil, status.Error(codes.InvalidArgument, "Missing ip")
			}

			ipNet := net.ParseIP(ip)

			_, cidrNet, err := net.ParseCIDR(config.TrustedSubnet)
			if err != nil {
				return nil, status.Error(codes.Internal, "Can't parse CIDR")
			}

			if !cidrNet.Contains(ipNet) {
				return nil, status.Error(codes.PermissionDenied, "Can't parse CIDR")
			}
		}

		return handler(ctx, req)
	}
}
