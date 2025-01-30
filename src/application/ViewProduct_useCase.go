package application

import (
	"demo/src/domain/entities"
	"demo/src/infraestructure/repositories"
)

type GetProducts struct {
	db repositories.ProductRepository
}

func NewGetProducts(db repositories.ProductRepository) *GetProducts {
	return &GetProducts{db: db}
}

func (gp *GetProducts) Execute() ([]entities.Product,error) { //diccionario
	res,err := gp.db.GetAll()
	if err != nil {
		return res,err
	}
	return res,nil
}
