package cart

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

func (ch *cartHandler) UpdateCartQuantity(ctx context.Context, request *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantityResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &cart.UpdateCartQuantityResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ch.cartSvc.UpdateCartQuantity(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
