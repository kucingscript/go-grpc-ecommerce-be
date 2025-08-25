package auth

import (
	"context"
	"errors"

	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"golang.org/x/crypto/bcrypt"
)

func (as *authService) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	if request.NewPassword != request.NewPasswordConfirmation {
		return &auth.ChangePasswordResponse{
			BaseResponse: utils.BadRequestResponse("New password and new password confirmation do not match"),
		}, nil
	}

	jwtToken, err := jwtModel.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := jwtModel.GetClaimsFromToken(jwtToken)
	if err != nil {
		return nil, err
	}

	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &auth.ChangePasswordResponse{
			BaseResponse: utils.BadRequestResponse("User not found"),
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return &auth.ChangePasswordResponse{
				BaseResponse: utils.BadRequestResponse("Old password is mismatched"),
			}, nil
		}

		return nil, err
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = as.authRepository.UpdateUserPassword(ctx, user.ID, string(hashedNewPassword), user.FullName)
	if err != nil {
		return nil, err
	}

	return &auth.ChangePasswordResponse{
		BaseResponse: utils.SuccessResponse("Change password successfully"),
	}, nil
}
