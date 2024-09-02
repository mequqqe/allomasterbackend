package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type BranchRepository struct {
	DB *gorm.DB
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{DB: db}
}

func (r *BranchRepository) AddBranch(branch models.Branch) (models.Branch, error) {
	err := r.DB.Create(&branch).Error
	return branch, err
}

func (r *BranchRepository) GetBranchesByCompanyId(companyId uint) ([]models.Branch, error) {
	var branches []models.Branch
	err := r.DB.Where("company_id = ?", companyId).Find(&branches).Error
	return branches, err
}
