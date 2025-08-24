package main

import (
	"context"
	"log"
	"net"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/config"
	authHandler "github.com/kucingscript/go-grpc-ecommerce-be/internal/handler/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/middleware"
	authRepository "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/auth"
	authService "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Panicf("failed to run server: %v", err)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	lis, err := net.Listen("tcp", cfg.GRPC_PORT)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	defer lis.Close()

	db := database.ConnectDB(ctx, cfg.DB_URI)
	log.Println("connected to database")

	defer db.Close()

	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository)
	authHandler := authHandler.NewAuthHandle(authService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ErrorMiddleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)

	if cfg.ENVIRONMENT == "dev" {
		reflection.Register(serv)
		log.Println("reflection is enabled")
	}

	log.Printf("server listening at %s", lis.Addr().String())
	if err := serv.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}

	return nil
}
