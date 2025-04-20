package service

import (
	"errors"
	"eventhub-backend/internal/repository"
)

type OrganizationService struct {
	organizationRepo repository.GormOrganizationRepository
}

func NewOrganizationService(organizationRepo repository.GormOrganizationRepository) *OrganizationService {
	return &OrganizationService{organizationRepo: organizationRepo}
}

func (s *OrganizationService) Create(name string, founderID uint) error {
	exists, err := s.organizationRepo.IsOrganizationExist(name)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("organization_already_exists")
	}

	return s.organizationRepo.Create(name, founderID)
}

func (s *OrganizationService) GetAll(userID uint) ([]repository.OrganizationModel, []repository.OrganizationModel, error) {
	return s.organizationRepo.GetAll(userID)
}

func (s *OrganizationService) JoinByCode(userID uint, code string) error {
	return s.organizationRepo.JoinByCode(userID, code)
}
