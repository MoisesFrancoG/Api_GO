package product

import (
	"demo/src/domain"
	"demo/src/domain/entities"
)

type GetProductById struct {
	productRepository domain.IProduct
}

func NewGetProductById(repo domain.IProduct) *GetProductById {
	return &GetProductById{productRepository: repo}
}

func (gp *GetProductById) Execute(id int) (*entities.Product, error) {
	return gp.productRepository.GetById(id)
}
