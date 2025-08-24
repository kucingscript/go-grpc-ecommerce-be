package jwt

import (
	"context"
	"strings"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"google.golang.org/grpc/metadata"
)

func ParseTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", utils.UnauthenticatedResponse()
	}

	bearerToken, ok := md["authorization"]
	if !ok {
		return "", utils.UnauthenticatedResponse()
	}

	if len(bearerToken) == 0 {
		return "", utils.UnauthenticatedResponse()
	}

	tokenSplit := strings.Split(bearerToken[0], " ")
	if len(tokenSplit) != 2 {
		return "", utils.UnauthenticatedResponse()
	}

	if tokenSplit[0] != "Bearer" {
		return "", utils.UnauthenticatedResponse()
	}

	return tokenSplit[1], nil
}
