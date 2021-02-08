package cartgrpc

import (
	i "github.com/barshat7/foodiecart/data/item"
	c "github.com/barshat7/foodiecart/data/cart"
	"context"
)

type Server struct {

}

func (s *Server) CreateCart(ctx context.Context, in *CreateCartRequest) (*CartID, error) {

	var items [] i.Item
	for _, cartItem := range in.CartItems {
		item := i.New(cartItem.ItemID, cartItem.ItemName, cartItem.ItemPrice)
		items = append(items, *item)
	}
	cart := c.New(in.UserID, items)
	persistedCartID, err := cart.Save()
	if (err != nil) {
		panic(err)
	}
	return &CartID {Value: persistedCartID}, nil
}