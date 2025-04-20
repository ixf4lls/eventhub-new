package repository

import (
	"errors"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) FindByUsername(username string) (*UserModel, error) {
	var user UserModel
	err := r.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, err
}

func (r *GormUserRepository) IsUsernameTaken(username string) (bool, error) {
	var user UserModel
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *GormUserRepository) Create(user *UserModel) error {
	return r.db.Create(user).Error
}
