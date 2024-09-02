package controllers

import (
	"encoding/json"
	"main/models"
	"main/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type CompanyController struct {
	CompanyService services.CompanyService
}

func (cc *CompanyController) Register(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registeredCompany, err := cc.CompanyService.Register(company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(registeredCompany)
}

func (cc *CompanyController) Login(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := cc.CompanyService.Login(request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (c *CompanyController) GetCompanyInfo(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)

	companyIdFloat, ok := claims["nameid"].(float64)
	if !ok {
		http.Error(w, "Invalid company ID", http.StatusUnauthorized)
		return
	}

	companyId := uint(companyIdFloat)

	company, err := c.CompanyService.GetCompanyById(companyId)
	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}
