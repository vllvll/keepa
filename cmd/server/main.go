package main

import (
	"github.com/vllvll/keepa/internal/handlers"
	"github.com/vllvll/keepa/internal/middlewares"
	"github.com/vllvll/keepa/internal/repositories"
	"github.com/vllvll/keepa/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	conf "github.com/vllvll/keepa/internal/config"
	"github.com/vllvll/keepa/pkg/postgres"

	// Импортируем пакет со сгенерированными protobuf-файлами
	pb "github.com/vllvll/keepa/proto"
)

func main() {
	config, err := conf.CreateServerConfig()
	if err != nil {
		log.Fatalf("Error with config: %v", err)
	}

	db, err := postgres.ConnectDatabase(config.DatabaseDsn)
	if err != nil {
		log.Fatalf("Error with database: %v", err)
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	bankCardRepository := repositories.NewBankCardRepository(db)
	textRepository := repositories.NewTextRepository(db)
	binaryRepository := repositories.NewBinaryRepository(db)
	cryptService := services.NewCrypt(config.Key)

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(middlewares.TrustSubnet(), middlewares.Auth(userRepository, tokenRepository)))

	go func() {
		listen, err := net.Listen("tcp", config.Address)
		if err != nil {
			log.Fatal(err)
		}

		pb.RegisterKeepaServer(s, handlers.NewHandler(
			userRepository,
			tokenRepository,
			bankCardRepository,
			textRepository,
			binaryRepository,
			cryptService,
		),
		)

		if err := s.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	for {
		select {
		case <-c:
			s.GracefulStop()
		}
	}
}
