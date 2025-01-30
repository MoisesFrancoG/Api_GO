package main

import (
	"demo/src/application"
	"demo/src/infraestructure"
	"demo/src/infraestructure/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	productRepo := repositories.NewProductRepository(mysql.DB)

	createProduct := application.NewCreateProduct(*productRepo)
	getProducts := application.NewGetProducts(*productRepo)
	updateProduct := application.NewUpdateProduct(*productRepo)

	
	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)

	router := gin.Default()
	productRoutes := infraestructure.NewProductRoutes(createProductController, getProductsController, updateProductController)
	productRoutes.SetupRoutes(router)

	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
