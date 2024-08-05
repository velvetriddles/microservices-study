package usecase

import (
	"product-service/internal/domain/model"
	"product-service/internal/repository"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) AddProduct(product *model.Product) error {
	return uc.repo.AddProduct(product)
}

func (uc *ProductUsecase) GetProduct(id int32) (*model.Product, error) {
	return uc.repo.GetProduct(id)
}

func (uc *ProductUsecase) ListProducts() ([]*model.Product, error) {
	return uc.repo.ListProducts()
}
