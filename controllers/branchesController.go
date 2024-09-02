package controllers

import (
	"encoding/json"
	"main/models"
	"main/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type BranchesController struct {
	BranchService *services.BranchService
}

func NewBranchesController(service *services.BranchService) *BranchesController {
	return &BranchesController{BranchService: service}
}

func (c *BranchesController) AddBranch(w http.ResponseWriter, r *http.Request) {
	var branch models.Branch
	err := json.NewDecoder(r.Body).Decode(&branch)
	if err != nil {
		http.Error(w, "Invalid branch data", http.StatusBadRequest)
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

	createdBranch, err := c.BranchService.AddBranch(branch, companyId)
	if err != nil {
		http.Error(w, "Could not add branch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBranch)
}

func (c *BranchesController) GetMyBranches(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)

	companyIdFloat, ok := claims["nameid"].(float64)
	if !ok {
		http.Error(w, "Invalid company ID", http.StatusUnauthorized)
		return
	}

	companyId := uint(companyIdFloat)

	branches, err := c.BranchService.GetBranchesByCompanyId(companyId)
	if err != nil {
		http.Error(w, "Could not retrieve branches", http.StatusInternalServerError)
		return
	}

	if len(branches) == 0 {
		http.Error(w, "No branches found for this company", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branches)
}
