package memory

import (
	"errors"
	"product-service/internal/domain/model"
	"product-service/internal/repository"
	"sync"
)

type ProductMemoryRepository struct {
	sync.RWMutex
	products map[int32]*model.Product
	nextID   int32
}

func NewProductMemoryRepository() repository.ProductRepository {
	return &ProductMemoryRepository{
		products: make(map[int32]*model.Product),
		nextID:   1,
	}
}

func (r *ProductMemoryRepository) AddProduct(product *model.Product) error {
	r.Lock()
	defer r.Unlock()

	product.ID = r.nextID
	r.products[r.nextID] = product
	r.nextID++

	return nil
}

func (r *ProductMemoryRepository) GetProduct(id int32) (*model.Product, error) {
	r.RLock()
	defer r.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (r *ProductMemoryRepository) ListProducts() ([]*model.Product, error) {
	r.RLock()
	defer r.RUnlock()

	var productList []*model.Product
	for _, product := range r.products {
		productList = append(productList, product)
	}

	return productList, nil
}
