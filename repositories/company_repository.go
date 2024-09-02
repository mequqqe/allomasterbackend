package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type CompanyRepositoryInterface interface {
	Create(company models.Company) (models.Company, error)
	GetById(id uint) (models.Company, error)
	GetByEmail(email string) (models.Company, error)
}

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

// Реализация интерфейса
func (r *CompanyRepository) Create(company models.Company) (models.Company, error) {
	err := r.DB.Create(&company).Error
	return company, err
}

func (r *CompanyRepository) GetById(id uint) (models.Company, error) {
	var company models.Company
	err := r.DB.First(&company, id).Error
	return company, err
}

func (r *CompanyRepository) GetByEmail(email string) (models.Company, error) {
	var company models.Company
	err := r.DB.Where("mail = ?", email).First(&company).Error
	return company, err
}
