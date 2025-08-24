package auth

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
)

func (a *authHandler) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &auth.LogoutResponse{
			BaseResponse: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := a.authService.Logout(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
