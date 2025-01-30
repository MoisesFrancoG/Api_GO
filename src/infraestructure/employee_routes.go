package infraestructure

import "github.com/gin-gonic/gin"

type EmployeeRoutes struct {
	CreateEmployeeController *CreateEmployeeController
	GetEmployeesController *GetEmployeesController
	UpdateEmployeeController *UpdateEmployeeController
	DeleteEmployeeController *DeleteEmployeeController
}

func NewEmployeeRoutes(createEmployeeController *CreateEmployeeController,getGetEmployeesController *GetEmployeesController,updateEmployeeController *UpdateEmployeeController, deleteEmployeeController *DeleteEmployeeController) *EmployeeRoutes {
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