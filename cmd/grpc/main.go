package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/config"
	authHandler "github.com/kucingscript/go-grpc-ecommerce-be/internal/handler/auth"
	cartHandler "github.com/kucingscript/go-grpc-ecommerce-be/internal/handler/cart"
	productHandler "github.com/kucingscript/go-grpc-ecommerce-be/internal/handler/product"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/middleware"
	authRepository "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/auth"
	cartRepository "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/cart"
	productRepository "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/product"
	authService "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/auth"
	cartService "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/cart"
	productService "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/product"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
	"github.com/kucingscript/go-grpc-ecommerce-be/pkg/database"
	gocache "github.com/patrickmn/go-cache"
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

	cacheService := gocache.New(time.Hour*24, time.Hour)
	authMiddleware := middleware.NewAuthMiddleware(cacheService)

	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository, cfg.JWT_SECRET, cacheService)
	authHandler := authHandler.NewAuthHandler(authService)

	productRepository := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepository, cfg.STORAGE_SERVICE_URL)
	productHandler := productHandler.NewProductHandler(productService)

	cartRepository := cartRepository.NewCartRepository(db)
	cartService := cartService.NewCartService(productRepository, cartRepository, cfg.STORAGE_SERVICE_URL)
	cartHandler := cartHandler.NewCartHandler(cartService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ErrorMiddleware,
			authMiddleware.Middleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)
	product.RegisterProductServiceServer(serv, productHandler)
	cart.RegisterCartServiceServer(serv, cartHandler)

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
