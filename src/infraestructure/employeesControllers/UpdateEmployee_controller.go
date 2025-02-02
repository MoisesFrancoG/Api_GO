package employeesControllers

import (
	"demo/src/application/employees"
	"demo/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEmployeeController struct {
	useCase_up *employees.UpdateEmployeeById
}

func NewUpdateEmployeeController(useCase_up *employees.UpdateEmployeeById) *UpdateEmployeeController {
	return &UpdateEmployeeController{useCase_up: useCase_up}
}

func (upe *UpdateEmployeeController) Execute(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var updatedEmployee entities.Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err = upe.useCase_up.Execute(id,&updatedEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "Employee": updatedEmployee})
}