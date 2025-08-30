package cart

import (
	"context"

	cartRepo "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/cart"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/product"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
	ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error)
	DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error)
	UpdateCartQuantity(ctx context.Context, request *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantityResponse, error)
}

type cartService struct {
	productRepository product.IProductRepository
	cartRepository    cartRepo.ICartRepository
	StorageServiceUrl string
}

func NewCartService(productRepository product.IProductRepository, cartRepository cartRepo.ICartRepository, storageServiceUrl string) ICartService {
	return &cartService{
		productRepository: productRepository,
		cartRepository:    cartRepository,
		StorageServiceUrl: storageServiceUrl,
	}
}
