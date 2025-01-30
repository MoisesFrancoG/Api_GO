package application

import (
	"demo/src/domain"
	"demo/src/domain/entities"
)

type CreateProduct struct {
	productRepository domain.IProduct
}

func NewCreateProduct(repo domain.IProduct) CreateProduct {
	return CreateProduct{productRepository: repo}
}

func (cp *CreateProduct) Execute(product *entities.Product) error {
	return cp.productRepository.SaveProductWithParams(product)
}
