package product

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (s *productService) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	request.Name = s.htmlSanitizer.Sanitize(request.Name)
	request.Description = s.htmlSanitizer.Sanitize(request.Description)

	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != model.USER_ROLE_ADMIN {
		return nil, utils.UnauthenticatedResponse()
	}

	imagePath := filepath.Join("storage", "product", request.ImageFileName)
	_, err = os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &product.CreateProductResponse{
				BaseResponse: utils.BadRequestResponse("Image file not found"),
			}, nil
		}

		return nil, err
	}

	now := time.Now()
	newProduct := model.Product{
		ID:            uuid.NewString(),
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		CreatedAt:     now,
		CreatedBy:     &claims.FullName,
	}

	err = s.productRepository.CreateNewProduct(ctx, &newProduct)
	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
		BaseResponse: utils.SuccessResponse("Create product successfully"),
		Id:           newProduct.ID,
	}, nil
}
