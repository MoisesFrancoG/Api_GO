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
	employeeRepo := repositories.NewEmployeeRepository(mysql.DB)

	createProduct := application.NewCreateProduct(productRepo)
	getProducts := application.NewGetProducts(productRepo)
	updateProduct := application.NewUpdateProduct(productRepo)
	deleteProduct := application.NewDeleteProduct(productRepo)

	createEmployee := application.NewCreateEmployee(employeeRepo)
	getEmployees := application.NewGetAllEmployees(employeeRepo)
	updateEmployee := application.NewUpdateEmployeeById(employeeRepo)
	deleteEmployee := application.NewDeleteEmployeeById(employeeRepo)

	
	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProduct)

	createEmployeeController := infraestructure.NewCreateEmployeeController(createEmployee)
	getEmployeesController := infraestructure.NewGetEmployeesController(getEmployees)
	updateEmployeeController := infraestructure.NewUpdateEmployeeController(updateEmployee)
	deleteEmployeeController := infraestructure.NewDeleteEmployeeController(deleteEmployee)

	router := gin.Default()
	productRoutes := infraestructure.NewProductRoutes(createProductController, getProductsController, updateProductController,deleteProductController)
	employeeRoutes := infraestructure.NewEmployeeRoutes(createEmployeeController, getEmployeesController,updateEmployeeController,deleteEmployeeController)
	productRoutes.SetupRoutes(router)
	employeeRoutes.SetupRoutes(router)

	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
