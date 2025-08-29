package product

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (ph *productHandler) DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &product.DeleteProductResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ph.productSvc.DeleteProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
