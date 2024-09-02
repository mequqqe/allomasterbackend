package services

import (
	"errors"
	"main/models"
	"main/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type CompanyService struct {
	Repo repositories.CompanyRepositoryInterface
}

func NewCompanyService(repo repositories.CompanyRepositoryInterface) CompanyService {
	return CompanyService{Repo: repo}
}

func (s *CompanyService) Register(company models.Company) (models.Company, error) {
	existingCompany, err := s.Repo.GetByEmail(company.Mail)
	if err == nil && existingCompany.ID != 0 {
		return models.Company{}, errors.New("Company already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Company{}, err
	}

	company.Password = string(hashedPassword)
	return s.Repo.Create(company)
}

func (s *CompanyService) Login(email, password string) (string, error) {
	company, err := s.Repo.GetByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(password)) != nil {
		return "", errors.New("Invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nameid":      company.ID,
		"CompanyName": company.CompanyName,
		"exp":         time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *CompanyService) GetCompanyById(id uint) (models.Company, error) {
	return s.Repo.GetById(id)
}
