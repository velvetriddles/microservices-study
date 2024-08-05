package memory

import (
	"order-service/internal/domain/model"
	"order-service/internal/repository"
	"sync"
)

type OrderMemoryRepository struct {
	sync.RWMutex
	orders map[int32]*model.Order
	nextID int32
}

func NewOrderMemoryRepository() repository.OrderRepository {
	return &OrderMemoryRepository{
		orders: make(map[int32]*model.Order),
		nextID: 1,
	}
}

func (r *OrderMemoryRepository) CreateOrder(order *model.Order) error {
	r.Lock()
	defer r.Unlock()

	r.orders[r.nextID] = order
	r.nextID++

	return nil
}
