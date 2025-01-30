package application

import (
	"demo/src/domain"
	"demo/src/domain/entities"
)

type CreateEmployee struct {
	employeeRepository domain.IEmployee
}

func NewCreateEmployee(employeeRepository domain.IEmployee) CreateEmployee {
	return CreateEmployee{employeeRepository: employeeRepository}
}

func (ce *CreateEmployee) Execute(employee *entities.Employee) error {
	return ce.employeeRepository.SaveEmployee(employee)
}
