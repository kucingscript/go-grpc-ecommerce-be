package product

import (
	"context"

	productRepo "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/product"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
	"github.com/microcosm-cc/bluemonday"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error)
	DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error)
	EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error)
	DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error)
	ListProduct(ctx context.Context, request *product.ListProductRequest) (*product.ListProductResponse, error)
	ListProductAdmin(ctx context.Context, request *product.ListProductAdminRequest) (*product.ListProductAdminResponse, error)
	HighlightProducts(ctx context.Context, request *product.HighlightProductsRequest) (*product.HighlightProductsResponse, error)
}

type productService struct {
	productRepository productRepo.IProductRepository
	htmlSanitizer     *bluemonday.Policy
	StorageServiceUrl string
}

func NewProductService(productRepository productRepo.IProductRepository, storageServiceUrl string) IProductService {
	sanitizer := bluemonday.UGCPolicy()

	return &productService{
		productRepository: productRepository,
		htmlSanitizer:     sanitizer,
		StorageServiceUrl: storageServiceUrl,
	}
}
