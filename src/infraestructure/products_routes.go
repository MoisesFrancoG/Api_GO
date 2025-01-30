package infraestructure

import (
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	CreateProductController *CreateProductController
	GetProductsController   *GetProductsController
	UpdateProductController *UpdateProductController
}

func NewProductRoutes(cpc *CreateProductController, gpc *GetProductsController,upc *UpdateProductController) *ProductRoutes {
	return &ProductRoutes{
		CreateProductController: cpc,
		GetProductsController: gpc,
		UpdateProductController: upc,
	}
}

func (pr *ProductRoutes) SetupRoutes(router *gin.Engine) {
	router.POST("/products", func(c *gin.Context) {
		pr.CreateProductController.Execute(c)
	})

	router.GET("/products", func(c *gin.Context) {
		pr.GetProductsController.Execute(c)
	})
	router.PUT("/products/:id", func(c *gin.Context) {
		pr.UpdateProductController.Execute(c)
	})
}
