package application

import (
	"demo/src/domain/entities"
	"demo/src/infraestructure/repositories"
)

type UpdateProduct struct {
	db repositories.ProductRepository
}

func NewUpdateProduct(db repositories.ProductRepository) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (up *UpdateProduct) Execute(id int, updatedProduct *entities.Product) error {
	err := up.db.EditById(id, updatedProduct)
	if err != nil {
		return err
	}
	return nil
}
