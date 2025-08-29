package product

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (ph *productHandler) HighlightProducts(ctx context.Context, request *product.HighlightProductsRequest) (*product.HighlightProductsResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &product.HighlightProductsResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ph.productSvc.HighlightProducts(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
