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
		item := i.New(cartItem.ItemID, cartItem.ItemName, cartItem.ItemPrice, cartItem.Quantity)
		items = append(items, *item)
	}
	cart := c.New(in.UserID, items)
	persistedCartID, err := cart.Save()
	if (err != nil) {
		panic(err)
	}
	return &CartID {Value: persistedCartID}, nil
}

func (s *Server) RetrieveCartByUserID(ctx context.Context, in *UserID) (*CartResponse, error) {
	userID := in.Value
	cart := c.FindActiveCartOfUser(userID)
	cartResponse := &CartResponse{Id: cart.ID, UserID: cart.UserID, CreatedDate: cart.CreatedDate.String()}
	for _, item := range cart.Items {
		cartItemResponse := CartItemResponse{Id: item.ID, CartID: item.CartID, ItemID: item.ItemID, ItemName: item.ItemName, ItemPrice: item.ItemPrice, Quantity: item.Quantity}
		cartResponse.CartItems = append(cartResponse.CartItems, &cartItemResponse)
	}

	return cartResponse, nil
}

func (s *Server) mustEmbedUnimplementedCartServiceServer() {
	// Do Nothing
}

func (s *Server) AddItemToCart(ctx context.Context, in *AddCartItemRequest) (*CartID, error) {
	var items [] i.Item
	for _, cartItem := range in.CartItems {
		item := i.New(cartItem.ItemID, cartItem.ItemName, cartItem.ItemPrice, cartItem.Quantity)
		items = append(items, *item)
	}
	cid := c.AddItemsToCartForUser(in.UserID, items)
	return &CartID {Value: cid}, nil
}