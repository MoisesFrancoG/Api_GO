package infraestructure

import (
	"demo/src/infraestructure/employeesControllers"

	"github.com/gin-gonic/gin"
)

type EmployeeRoutes struct {
	CreateEmployeeController *employeesControllers.CreateEmployeeController
	GetEmployeesController *employeesControllers.GetEmployeesController
	UpdateEmployeeController *employeesControllers.UpdateEmployeeController
	DeleteEmployeeController *employeesControllers.DeleteEmployeeController
}

func NewEmployeeRoutes(createEmployeeController *employeesControllers.CreateEmployeeController,getGetEmployeesController *employeesControllers.GetEmployeesController,updateEmployeeController *employeesControllers.UpdateEmployeeController, deleteEmployeeController *employeesControllers.DeleteEmployeeController) *EmployeeRoutes {
	return &EmployeeRoutes{
		CreateEmployeeController: createEmployeeController,
		GetEmployeesController: getGetEmployeesController,
		UpdateEmployeeController: updateEmployeeController,
		DeleteEmployeeController: deleteEmployeeController,
	}
}

func (em *EmployeeRoutes) SetupRoutes(router *gin.Engine) {
	router.POST("/employees", func(c *gin.Context) {
		em.CreateEmployeeController.Execute(c)
	})

	router.GET("/employees", func(c *gin.Context) {
		em.GetEmployeesController.Execute(c)
	})

	router.PUT("/employees/:id",func(c *gin.Context){
		em.UpdateEmployeeController.Execute(c)
	})

	router.DELETE("employees/:id", func(ctx *gin.Context) {
		em.DeleteEmployeeController.Execute(ctx)
	})
}