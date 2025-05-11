package domain

import "time"

type Event struct {
	ID          uint
	Title       string
	Description string
	Category    string
	IsPublic    bool
	Status      string
	Date        string
	StartTime   string
	EndTime     string
	Location    string
	CreatorId   uint
}

type CreateEventInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Location    string    `json:"location"`
	IsPublic    bool      `json:"is_public"`
	Date        time.Time `gorm:"type:date" json:"date"`
	StartTime   string    `gorm:"type:time" json:"start_time"`
	EndTime     string    `gorm:"type:time" json:"end_time"`
}
