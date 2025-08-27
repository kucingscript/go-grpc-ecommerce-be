package product

import (
	productSvc "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/product"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

type productHandler struct {
	product.UnimplementedProductServiceServer
	productSvc productSvc.IProductService
}

func NewProductHandler(productSvc productSvc.IProductService) *productHandler {
	return &productHandler{
		productSvc: productSvc,
	}
}
