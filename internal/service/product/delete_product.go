package product

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (s *productService) DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != model.USER_ROLE_ADMIN {
		return nil, utils.UnauthenticatedResponse()
	}

	productExist, err := s.productRepository.GetProductByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if productExist == nil {
		return &product.DeleteProductResponse{
			BaseResponse: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	err = s.productRepository.DeleteProduct(ctx, request.Id, time.Now(), claims.FullName)
	if err != nil {
		return nil, err
	}

	imagePath := filepath.Join("storage", "product", productExist.ImageFileName)
	err = os.Remove(imagePath)
	if err != nil {
		return nil, err
	}

	return &product.DeleteProductResponse{
		BaseResponse: utils.SuccessResponse("Delete product successfully"),
	}, nil
}
