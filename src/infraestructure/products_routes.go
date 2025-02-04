package infraestructure

import (
	"demo/src/infraestructure/products"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	CreateProductController *products.CreateProductController
	GetProductsController   *products.GetProductsController
	UpdateProductController *products.UpdateProductController
	DeleteProductController *products.DeleteProductController
	GetProductById *products.GetProductByIdController
}

func NewProductRoutes(cpc *products.CreateProductController, gpc *products.GetProductsController,upc *products.UpdateProductController,dpc *products.DeleteProductController,gbd *products.GetProductByIdController) *ProductRoutes {
	return &ProductRoutes{
		CreateProductController: cpc,
		GetProductsController: gpc,
		UpdateProductController: upc,
		DeleteProductController: dpc,
		GetProductById: gbd,
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

	router.GET("/products/:id", func(c *gin.Context) { 
		pr.GetProductById.Execute(c)
	})
}
