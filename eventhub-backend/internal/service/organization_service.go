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
	exists, err := s.organizationRepo.IsNameTaken(name)
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

func (s *OrganizationService) GetEvents(orgID uint) ([]repository.EventResponse, []repository.EventResponse, error) {
	exists, err := s.organizationRepo.IsOrganizationExist(orgID)
	if err != nil {
		return nil, nil, err
	}
	if !exists {
		return nil, nil, errors.New("organization not exists")
	}

	return s.organizationRepo.GetEvents(orgID)
}

func (s *OrganizationService) GetCreator(orgID uint) (uint, error) {
	return s.organizationRepo.GetCreator(orgID)
}

func (s *OrganizationService) GetByID(orgID, userID uint) (repository.OrganizationModel, bool, error) {
	exists, err := s.organizationRepo.IsOrganizationExist(orgID)
	if err != nil {
		return repository.OrganizationModel{}, false, err
	}
	if !exists {
		return repository.OrganizationModel{}, false, errors.New("organization not exists")
	}

	return s.organizationRepo.GetByID(orgID, userID)
}

func (s *OrganizationService) GetMembers(orgID uint) ([]repository.UserAsMember, error) {
	exists, err := s.organizationRepo.IsOrganizationExist(orgID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("organization not exists")
	}

	return s.organizationRepo.GetMembers(orgID)
}
