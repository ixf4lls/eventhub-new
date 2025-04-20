package service

import (
	"errors"
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	userRepo repository.UserRepository
}

func NewRegisterService(userRepo repository.UserRepository) *RegisterService {
	return &RegisterService{userRepo: userRepo}
}

func (s *RegisterService) Register(input domain.RegisterInput) error {
	usernameTaken, err := s.userRepo.IsUsernameTaken(input.Username)
	if err != nil {
		return err
	}
	if usernameTaken {
		return errors.New("username_taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &repository.UserModel{
		Username:     input.Username,
		PasswordHash: string(hash),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
	}

	return s.userRepo.Create(newUser)
}
