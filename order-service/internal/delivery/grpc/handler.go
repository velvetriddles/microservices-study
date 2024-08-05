package grpcorder

import (
	"context"

	"order-service/internal/usecase"
	pb "order-service/pkg/api/v1"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(uc *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: uc}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &usecase.Order{
		UserID:    req.UserId,
		ProductID: req.ProductId,
		Quantity:  req.Quantity,
	}

	err := h.usecase.CreateOrder(order)
	if err != nil {
		return &pb.CreateOrderResponse{Success: false, Error: err.Error()}, nil
	}

	return &pb.CreateOrderResponse{Success: true}, nil
}
