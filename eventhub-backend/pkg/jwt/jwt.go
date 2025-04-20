package customJwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtManager struct {
	secretKey string
}

func NewJwtManager(secretKey string) *JwtManager {
	return &JwtManager{secretKey: secretKey}
}

func (jm *JwtManager) GenerateAccessToken(userID uint) (string, error) {
	// key := []byte(config.Load().JwtSecretKey)

	claims := &jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Minute * 3).Unix(),
	}

	// refreshClaims := &jwt.MapClaims{
	// 	"sub": userId,
	// 	"exp": time.Now().Add(time.Hour * 24).Unix(),
	// 	"typ": "refresh",
	// 	"jti": uuid.NewString(),
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	return token.SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) GenerateRefreshToken(userID uint) (string, error) {
	claims := &jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"typ": "refresh",
		"jti": uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) ParseToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jm.secretKey), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid userID")
	}

	return uint(sub), nil
}
