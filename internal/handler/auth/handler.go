package auth

import (
	authSvc "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
)

type authHandler struct {
	auth.UnimplementedAuthServiceServer
	authService authSvc.IAuthService
}

func NewAuthHandle(authService authSvc.IAuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}
