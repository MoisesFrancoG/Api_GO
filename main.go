package main

import (
	"demo/src/application/employees"
	"demo/src/application/product"
	"demo/src/infraestructure"
	"demo/src/infraestructure/employeesControllers"
	"demo/src/infraestructure/products"
	"demo/src/infraestructure/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	productRepo := repositories.NewProductRepository(mysql.DB)
	employeeRepo := repositories.NewEmployeeRepository(mysql.DB)

	createProduct := product.NewCreateProduct(productRepo)
	getProducts := product.NewGetProducts(productRepo)
	updateProduct := product.NewUpdateProduct(productRepo)
	deleteProduct := product.NewDeleteProduct(productRepo)
	getById := product.NewGetProductById(productRepo)

	createEmployee := employees.NewCreateEmployee(employeeRepo)
	getEmployees := employees.NewGetAllEmployees(employeeRepo)
	updateEmployee := employees.NewUpdateEmployeeById(employeeRepo)
	deleteEmployee := employees.NewDeleteEmployeeById(employeeRepo)

	
	createProductController := products.NewCreateProductController(createProduct)
	getProductsController := products.NewGetProductsController(getProducts)
	updateProductController := products.NewUpdateProductController(updateProduct)
	deleteProductController := products.NewDeleteProductController(deleteProduct)
	getByIdController := products.NewGetProductByIdController(getById)

	createEmployeeController := employeesControllers.NewCreateEmployeeController(createEmployee)
	getEmployeesController := employeesControllers.NewGetEmployeesController(getEmployees)
	updateEmployeeController := employeesControllers.NewUpdateEmployeeController(updateEmployee)
	deleteEmployeeController := employeesControllers.NewDeleteEmployeeController(deleteEmployee)

	router := gin.Default()
	productRoutes := infraestructure.NewProductRoutes(createProductController, getProductsController, updateProductController,deleteProductController,getByIdController)
	employeeRoutes := infraestructure.NewEmployeeRoutes(createEmployeeController, getEmployeesController,updateEmployeeController,deleteEmployeeController)
	productRoutes.SetupRoutes(router)
	employeeRoutes.SetupRoutes(router)

	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
