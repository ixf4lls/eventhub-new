package repository

import "time"

type NotificationModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	EventID   uint      `json:"event_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func (NotificationModel) TableName() string {
	return "notifications"
}
