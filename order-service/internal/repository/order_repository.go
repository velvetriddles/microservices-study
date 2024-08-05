package repository

import "order-service/internal/domain/model"

type OrderRepository interface {
	CreateOrder(order *model.Order) error
}
