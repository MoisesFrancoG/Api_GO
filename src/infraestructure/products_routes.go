package infraestructure

import (
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	CreateProductController *CreateProductController
	GetProductsController   *GetProductsController
	UpdateProductController *UpdateProductController
	DeleteProductController *DeleteProductController
}

func NewProductRoutes(cpc *CreateProductController, gpc *GetProductsController,upc *UpdateProductController,dpc *DeleteProductController) *ProductRoutes {
	return &ProductRoutes{
		CreateProductController: cpc,
		GetProductsController: gpc,
		UpdateProductController: upc,
		DeleteProductController: dpc,
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
	router.DELETE("/products/:id",func(c *gin.Context) {
		pr.DeleteProductController.Execute(c)
	})
}
