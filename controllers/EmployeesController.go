package controllers

import (
	"encoding/json"
	"main/models"
	"main/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type EmployeesController struct {
	EmployeeService *services.EmployeeService
}

func NewEmployeesController(service *services.EmployeeService) *EmployeesController {
	return &EmployeesController{EmployeeService: service}
}

func (c *EmployeesController) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid employee data", http.StatusBadRequest)
		return
	}

	claims := r.Context().Value("claims").(jwt.MapClaims)

	// Предполагаем, что nameid хранится как float64
	companyIdFloat, ok := claims["nameid"].(float64)
	if !ok {
		http.Error(w, "Invalid company ID", http.StatusUnauthorized)
		return
	}

	// Преобразуем float64 в uint
	companyId := uint(companyIdFloat)

	// Присваиваем companyId сотруднику
	employee.CompanyID = companyId

	createdEmployee, err := c.EmployeeService.AddEmployee(employee)
	if err != nil {
		http.Error(w, "Could not add employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmployee)
}

func (c *EmployeesController) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := c.EmployeeService.GetAllEmployees()
	if err != nil {
		http.Error(w, "Failed to fetch employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
