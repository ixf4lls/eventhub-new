package repository

import "time"

type EventModel struct {
	ID             uint      `gorm:"primaryKey"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	IsPublic       bool      `json:"is_public"`
	Status         string    `json:"status"`
	Date           time.Time `json:"date"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	Location       string    `json:"location"`
	CreatorId      uint      `json:"creator_id"`
	OrganizationId uint      `json:"organization_id"`
}

type EventResponse struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Category       string `json:"category"`
	IsPublic       bool   `json:"is_public"`
	Status         string `json:"status"`
	Date           string `json:"date"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Location       string `json:"location"`
	CreatorId      uint   `json:"creator_id"`
	OrganizationId uint   `json:"organization_id"`
}

func (EventModel) TableName() string {
	return "events"
}

type EventParticipantModel struct {
	UserID  uint `gorm:"primaryKey" json:"user_id"`
	EventID uint `gorm:"primaryKey" json:"event_id"`
}

func (EventParticipantModel) TableName() string {
	return "event_participants"
}
