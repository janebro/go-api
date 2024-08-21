package service

import (
	"go-api/model"
	"go-api/repository"
)

type ProductService struct {
	repository *repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (pu *ProductService) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	productId, error := pu.repository.CreateProduct(product)

	if error != nil {
		return model.Product{}, error
	}

	product.Id = productId
	return product, nil
}

func (pu *ProductService) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
