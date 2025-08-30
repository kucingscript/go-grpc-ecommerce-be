package cart

import (
	"context"
	"time"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

func (s *cartService) UpdateCartQuantity(ctx context.Context, request *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantityResponse, error) {
	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != model.USER_ROLE_CUSTOMER && claims.Role != model.USER_ROLE_ADMIN {
		return nil, utils.UnauthenticatedResponse()
	}

	cartExists, err := s.cartRepository.GetCartByID(ctx, request.CartId)
	if err != nil {
		return nil, err
	}

	if cartExists == nil {
		return &cart.UpdateCartQuantityResponse{
			BaseResponse: utils.NotFoundResponse("Cart not found"),
		}, nil
	}

	if cartExists.UserID != claims.Subject {
		return &cart.UpdateCartQuantityResponse{
			BaseResponse: utils.BadRequestResponse("Cart not found"),
		}, nil
	}

	if request.NewQuantity == 0 {
		err = s.cartRepository.DeleteCart(ctx, request.CartId)
		if err != nil {
			return nil, err
		}

		return &cart.UpdateCartQuantityResponse{
			BaseResponse: utils.SuccessResponse("Update cart quantity successfully"),
		}, nil
	}

	now := time.Now()
	cartExists.Quantity = int64(request.NewQuantity)
	cartExists.UpdatedAt = &now
	cartExists.UpdatedBy = &claims.FullName

	err = s.cartRepository.UpdateCart(ctx, cartExists)
	if err != nil {
		return nil, err
	}

	return &cart.UpdateCartQuantityResponse{
		BaseResponse: utils.SuccessResponse("Update cart quantity successfully"),
	}, nil
}
