package product

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (ph *productHandler) ListProduct(ctx context.Context, request *product.ListProductRequest) (*product.ListProductResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &product.ListProductResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ph.productSvc.ListProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
