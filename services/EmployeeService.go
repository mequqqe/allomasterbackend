package services

import (
	"main/models"
	"main/repositories"
)

type EmployeeService struct {
	Repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) AddEmployee(employee models.Employee) (models.Employee, error) {
	// Пример: можно добавить шифрование пароля перед сохранением
	// employee.Password = HashPassword(employee.Password)
	return s.Repo.AddEmployee(employee)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAllEmployees()
}
