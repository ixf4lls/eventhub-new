package repository

import (
	"crypto/rand"
	"errors"

	"gorm.io/gorm"
)

type GormOrganizationRepository struct {
	db *gorm.DB
}

func NewGormOrganizationRepository(db *gorm.DB) *GormOrganizationRepository {
	return &GormOrganizationRepository{db: db}
}

func (r *GormOrganizationRepository) Create(name string, founderID uint) error {
	var inviteCode string
	var err error

	for {
		inviteCode, err = generateInviteCode()
		if err != nil {
			return err
		}

		var count int64
		if err := r.db.Model(&OrganizationModel{}).Where("invite_code = ?", inviteCode).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			break
		}
	}

	organization := OrganizationModel{
		Name:       name,
		FounderID:  founderID,
		InviteCode: inviteCode,
	}

	if err := r.db.Create(&organization).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormOrganizationRepository) IsNameTaken(name string) (bool, error) {
	var exists OrganizationModel
	err := r.db.Where("name = ?", name).First(&exists).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *GormOrganizationRepository) GetAll(userID uint) ([]OrganizationModel, []OrganizationModel, error) {
	var joinedIDs []uint
	if err := r.db.Table("organization_members").Where("user_id = ?", userID).Pluck("organization_id", &joinedIDs).Error; err != nil {
		return nil, nil, err
	}

	var joined []OrganizationModel
	if err := r.db.Where("id IN (?)", joinedIDs).Find(&joined).Error; err != nil {
		return nil, nil, err
	}

	var founded []OrganizationModel
	if err := r.db.Where("founder_id = ?", userID).Find(&founded).Error; err != nil {
		return nil, nil, err
	}

	return joined, founded, nil
}

func (r *GormOrganizationRepository) JoinByCode(userID uint, code string) error {
	var organization OrganizationModel
	if err := r.db.Where("invite_code = ?", code).First(&organization).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("organization not found")
		}
		return err
	}

	if userID == organization.FounderID {
		return errors.New("user is a creator of this organization")
	}

	var existingMember OrganizationMemberModel
	if err := r.db.Where("user_id = ? AND organization_id = ?", userID, organization.ID).First(&existingMember).Error; err == nil {
		return errors.New("user already a member of the organization")
	}

	orgMember := OrganizationMemberModel{
		UserID:         userID,
		OrganizationID: organization.ID,
	}

	if err := r.db.Create(&orgMember).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormOrganizationRepository) GetCreator(orgID uint) (uint, error) {
	var organization OrganizationModel
	err := r.db.Where("id = ?", orgID).First(&organization).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("organization not found")
		}
		return 0, err
	}

	return organization.FounderID, nil
}

func (r *GormOrganizationRepository) GetEvents(orgID uint) ([]EventResponse, []EventResponse, error) {
	var events []EventModel
	if err := r.db.Where("organization_id = ?", orgID).Find(&events).Error; err != nil {
		return nil, nil, err
	}

	var activeEvents, completedEvents []EventModel
	for _, event := range events {
		if event.Status == "active" {
			activeEvents = append(activeEvents, event)
		} else if event.Status == "completed" {
			completedEvents = append(completedEvents, event)
		}
	}

	return parseEventTime(activeEvents), parseEventTime(completedEvents), nil
}

func (r *GormOrganizationRepository) IsOrganizationExist(orgID uint) (bool, error) {
	var exists OrganizationModel
	err := r.db.Where("id = ?", orgID).First(&exists).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("organization not found")
		}
		return false, err
	}

	return true, nil
}

func (r *GormOrganizationRepository) GetByID(orgID, userID uint) (OrganizationModel, bool, error) {
	var organization OrganizationModel
	if err := r.db.Where("id = ?", orgID).First(&organization).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return OrganizationModel{}, false, errors.New("organization not found")
		}
		return OrganizationModel{}, false, err
	}

	return organization, userID == organization.FounderID, nil
}

func (r *GormOrganizationRepository) GetMembers(orgID uint) ([]UserAsMember, error) {
	var users []UserModel
	var usersResponse []UserAsMember
	var userIDs []uint

	if err := r.db.Table("organization_members").Where("organization_id = ?", orgID).Pluck("user_id", &userIDs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("участники не найдены")
		}
		return nil, err
	}

	if err := r.db.Table("users").Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("участники не найдены")
		}
		return nil, err
	}

	for _, user := range users {
		usersResponse = append(usersResponse, UserAsMember{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
		})
	}

	return usersResponse, nil
}

func generateInviteCode() (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(bytes); i++ {
		bytes[i] = charset[int(bytes[i])%len(charset)]
	}

	return string(bytes), nil
}
