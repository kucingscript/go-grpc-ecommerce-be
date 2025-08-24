package jwt

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
)

type JwtModelContextKey string

var JwtModelContextKeyValue JwtModelContextKey = "jwtModel"

type JwtClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (jc *JwtClaims) SetToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, JwtModelContextKeyValue, jc)
}

func GetClaimsFromToken(token string) (*JwtClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !tokenClaims.Valid {
		return nil, utils.UnauthenticatedResponse()
	}

	claims, ok := tokenClaims.Claims.(*JwtClaims)
	if ok {
		return claims, nil
	}

	return nil, utils.UnauthenticatedResponse()
}
