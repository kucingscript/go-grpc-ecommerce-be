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

func (s *productService) EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error) {
	request.Name = s.htmlSanitizer.Sanitize(request.Name)
	request.Description = s.htmlSanitizer.Sanitize(request.Description)

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
		return &product.EditProductResponse{
			BaseResponse: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	if productExist.ImageFileName != request.ImageFileName {
		newImagePath := filepath.Join("storage", "product", request.ImageFileName)
		_, err = os.Stat(newImagePath)
		if err != nil {
			if os.IsNotExist(err) {
				return &product.EditProductResponse{
					BaseResponse: utils.BadRequestResponse("Image file not found"),
				}, nil
			}

			return nil, err
		}

		oldImagePath := filepath.Join("storage", "product", productExist.ImageFileName)
		err = os.Remove(oldImagePath)
		if err != nil {
			return nil, err
		}
	}

	now := time.Now()
	updatedProduct := model.Product{
		ID:            request.Id,
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		UpdatedAt:     &now,
		UpdatedBy:     &claims.FullName,
	}

	err = s.productRepository.UpdateProduct(ctx, &updatedProduct)
	if err != nil {
		return nil, err
	}

	return &product.EditProductResponse{
		BaseResponse: utils.SuccessResponse("Edit product successfully"),
		Id:           request.Id,
	}, nil
}
