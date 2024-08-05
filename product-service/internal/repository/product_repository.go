package repository

import "product-service/internal/domain/model"

type ProductRepository interface {
	AddProduct(product *model.Product) error
	GetProduct(id int32) (*model.Product, error)
	ListProducts() ([]*model.Product, error)
}
