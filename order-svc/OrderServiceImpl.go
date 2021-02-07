package main

import (

	o "github.com/barshat7/foodieorders/service"
	d "github.com/barshat7/foodieorders/data"
	"context"
	_ "github.com/gofrs/uuid"
	"strconv"
	"fmt"
)

type server struct {
}

func (s *server) CreateOrder(ctx context.Context, in *o.OrderRequest) (*o.OrderID, error) {
	id, err := d.SaveOrder(in.UserID, in.TotalPrice)
	fmt.Println("Order Saved")
	return &o.OrderID{Value:strconv.Itoa(int(id))}, err
}

func (s *server) RetrieveOrder(ctx context.Context, in *o.OrderID) (*o.Order, error) {
	id := in.Value
	intID, _ := strconv.ParseInt(id, 10, 64)
	order := d.FindOrderByID(intID)

	return &o.Order{UserID: order.UserID, TotalPrice: order.TotalPrice, OrderDate: order.CreatedDate.String()}, nil
}

func (s *server) RetrieveOrdersByUserID(in *o.UserID, stream o.OrderService_RetrieveOrdersByUserIDServer) error {
	ordersByUserID := d.FindOrdersByUserID(in.Value)
	fmt.Println(len(ordersByUserID))
	for _,order := range ordersByUserID {
		oo := &o.Order{UserID: order.UserID, TotalPrice: order.TotalPrice, OrderDate: order.CreatedDate.String()}
		err := stream.Send(oo)
		if err != nil {
			return fmt.Errorf(
				"error sending message to stream : %v",
					err)
		}
	}
	return nil
}