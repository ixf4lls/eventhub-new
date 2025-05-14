package service

import "eventhub-backend/internal/repository"

type UserService struct {
	userRepo repository.GormUserRepository
}

func NewUserService(userRepo repository.GormUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByID(ID uint) (repository.UserResponse, error) {
	return s.userRepo.GetByID(ID)
}
