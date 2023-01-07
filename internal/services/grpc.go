package services

import (
	"context"
	pb "github.com/vllvll/keepa/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"sync"
)

var (
	ctx    context.Context
	client GRPCSender
	once   sync.Once
)

type GRPCSender struct {
	Client pb.KeepaClient
}

func NewGRPCSendClient() (*GRPCSender, error) {
	// устанавливаем соединение с сервером
	once.Do(func() {
		conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}

		c := pb.NewKeepaClient(conn)

		client = GRPCSender{
			Client: c,
		}
	})

	return &client, nil
}

func (g *GRPCSender) SetToken(token string) {
	md := metadata.New(map[string]string{"keepa-auth": token})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
}

func (g *GRPCSender) GetContext() context.Context {
	return ctx
}
