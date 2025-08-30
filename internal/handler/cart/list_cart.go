package cart

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

func (ch *cartHandler) ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &cart.ListCartResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ch.cartSvc.ListCart(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
