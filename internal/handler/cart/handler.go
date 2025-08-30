package cart

import (
	cartSvc "github.com/kucingscript/go-grpc-ecommerce-be/internal/service/cart"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

type cartHandler struct {
	cart.UnimplementedCartServiceServer
	cartSvc cartSvc.ICartService
}

func NewCartHandler(cartSvc cartSvc.ICartService) *cartHandler {
	return &cartHandler{
		cartSvc: cartSvc,
	}
}
