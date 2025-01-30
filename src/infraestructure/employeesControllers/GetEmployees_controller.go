package employeesControllers

import (
	"demo/src/application/employees"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetEmployeesController struct {
	useCase_ge *employees.GetAllEmployees
}

func NewGetEmployeesController(useCase_ge *employees.GetAllEmployees) *GetEmployeesController {
	return &GetEmployeesController{useCase_ge: useCase_ge}
}

func (gec *GetEmployeesController) Execute(c *gin.Context) {
	employees, err := gec.useCase_ge.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employees": employees})
}