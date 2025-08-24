package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"golang.org/x/crypto/bcrypt"
)

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			BaseResponse: utils.BadRequestResponse("Password and password confirmation do not match"),
		}, nil
	}

	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return &auth.RegisterResponse{
			BaseResponse: utils.BadRequestResponse("User already exists"),
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		ID:        uuid.NewString(),
		FullName:  request.FullName,
		Email:     request.Email,
		Password:  string(hashedPassword),
		RoleCode:  model.USER_ROLE_CUSTOMER,
		CreatedBy: &request.FullName,
	}

	err = as.authRepository.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		BaseResponse: utils.SuccessResponse("User registered successfully"),
	}, nil
}
