package employees

import (
	"demo/src/domain"
	"demo/src/domain/entities"
)

type GetAllEmployees struct {
	employeeRepository domain.IEmployee
}

func NewGetAllEmployees(employeeRepository domain.IEmployee) *GetAllEmployees {
	return &GetAllEmployees{employeeRepository: employeeRepository}
}

func (gae *GetAllEmployees) Execute() ([]entities.Employee, error) {
	return gae.employeeRepository.GetAllEmployees()
}
