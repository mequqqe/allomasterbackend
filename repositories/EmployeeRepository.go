package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) AddEmployee(employee models.Employee) (models.Employee, error) {
	err := r.DB.Create(&employee).Error
	return employee, err
}

func (r *EmployeeRepository) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Find(&employees).Error
	return employees, err
}
