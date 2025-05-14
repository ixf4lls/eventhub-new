package repository

import (
	"errors"

	"gorm.io/gorm"
)

type GormNotificationRepository struct {
	db *gorm.DB
}

func NewGormNotificationRepository(db *gorm.DB) *GormNotificationRepository {
	return &GormNotificationRepository{db: db}
}

func (r *GormNotificationRepository) Create(userID, eventID uint, msgType string) error {
	notification := NotificationModel{
		UserID:  userID,
		EventID: eventID,
		Type:    msgType,
	}

	if err := r.db.Create(&notification).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormNotificationRepository) GetAll(userID uint) ([]NotificationModel, error) {
	var notifications []NotificationModel

	if err := r.db.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("notifications not found")
		}

		return nil, err
	}

	return notifications, nil
}

func (r *GormNotificationRepository) Exists(eventID, userID uint, msgType string) (bool, error) {
	var notification NotificationModel

	if err := r.db.Where("user_id = ? AND event_id = ? AND type = ?", userID, eventID, msgType).First(&notification).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *GormNotificationRepository) GetById(notificationID uint) (NotificationModel, error) {
	var notification NotificationModel

	if err := r.db.Where("id = ?", notificationID).First(&notification).Error; err != nil {
		return NotificationModel{}, err
	}

	return notification, nil
}
