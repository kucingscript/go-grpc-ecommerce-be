package cart

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

func (s *cartService) DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
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
		return &cart.DeleteCartResponse{
			BaseResponse: utils.NotFoundResponse("Cart not found"),
		}, nil
	}

	if cartExists.UserID != claims.Subject {
		return &cart.DeleteCartResponse{
			BaseResponse: utils.BadRequestResponse("Cart not found"),
		}, nil
	}

	err = s.cartRepository.DeleteCart(ctx, request.CartId)
	if err != nil {
		return nil, err
	}

	return &cart.DeleteCartResponse{
		BaseResponse: utils.SuccessResponse("Success delete cart"),
	}, nil
}
