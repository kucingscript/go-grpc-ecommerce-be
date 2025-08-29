package product

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (ph *productHandler) EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &product.EditProductResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ph.productSvc.EditProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
