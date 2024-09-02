package services

import (
	"main/models"
	"main/repositories"
)

type BranchService struct {
	Repo *repositories.BranchRepository
}

func NewBranchService(repo *repositories.BranchRepository) *BranchService {
	return &BranchService{Repo: repo}
}

func (s *BranchService) AddBranch(branch models.Branch, companyId uint) (models.Branch, error) {
	branch.CompanyID = companyId
	return s.Repo.AddBranch(branch)
}

func (s *BranchService) GetBranchesByCompanyId(companyId uint) ([]models.Branch, error) {
	return s.Repo.GetBranchesByCompanyId(companyId)
}
