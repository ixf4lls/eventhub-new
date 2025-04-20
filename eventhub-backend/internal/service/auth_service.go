package service

import (
	"errors"
	"eventhub-backend/internal/config"
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/repository"
	customHash "eventhub-backend/pkg/hash"
	customJwt "eventhub-backend/pkg/jwt"
)

type AuthService struct {
	cfg        config.Config
	userRepo   repository.UserRepository
	jwtManager customJwt.Manager
}

func NewAuthService(cfg config.Config, userRepo repository.UserRepository, jwtManager customJwt.Manager) *AuthService {
	return &AuthService{cfg: cfg, userRepo: userRepo, jwtManager: jwtManager}
}

func (s *AuthService) RefreshTokens(refreshToken string) (string, string, error) {
	userID, _ := s.jwtManager.ParseToken(refreshToken)
	access, err := s.jwtManager.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refresh, err := s.jwtManager.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *AuthService) Login(input domain.LoginInput) (string, string, error) {
	user, err := s.userRepo.FindByUsername(input.Username)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !customHash.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
