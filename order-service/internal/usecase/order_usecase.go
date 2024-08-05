package usecase

import (
	"context"
	"errors"
	"order-service/internal/domain/model"
	"order-service/internal/repository"

	productpb "product-service/pkg/api/v1"

	"google.golang.org/grpc"
)

type Order struct {
	UserID    string
	ProductID int32
	Quantity  int32
}

type OrderUsecase struct {
	repo        repository.OrderRepository
	productConn *grpc.ClientConn
}

func NewOrderUsecase(repo repository.OrderRepository, productConn *grpc.ClientConn) *OrderUsecase {
	return &OrderUsecase{repo: repo, productConn: productConn}
}

func (uc *OrderUsecase) CreateOrder(order *Order) error {
	productClient := productpb.NewProductServiceClient(uc.productConn)

	productResp, err := productClient.GetProduct(context.Background(), &productpb.GetProductRequest{Id: order.ProductID})
	if err != nil {
		return errors.New("failed to get product information")
	}

	if productResp.Quantity < order.Quantity {
		return errors.New("not enough products in stock")
	}

	return uc.repo.CreateOrder(&model.Order{
		UserID:    order.UserID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
	})
}
