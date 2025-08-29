package product

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (ph *productHandler) ListProductAdmin(ctx context.Context, request *product.ListProductAdminRequest) (*product.ListProductAdminResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &product.ListProductAdminResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ph.productSvc.ListProductAdmin(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
