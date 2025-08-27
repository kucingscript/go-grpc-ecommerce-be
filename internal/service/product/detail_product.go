package product

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (s *productService) DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error) {
	productExist, err := s.productRepository.GetProductByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if productExist == nil {
		return &product.DetailProductResponse{
			BaseResponse: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	return &product.DetailProductResponse{
		BaseResponse: utils.SuccessResponse("Get product successfully"),
		Id:           productExist.ID,
		Name:         productExist.Name,
		Description:  productExist.Description,
		Price:        productExist.Price,
		ImageUrl:     fmt.Sprintf("%s/product/%s", s.StorageServiceUrl, productExist.ImageFileName),
	}, nil
}
