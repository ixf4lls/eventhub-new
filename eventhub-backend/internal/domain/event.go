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
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	Location       string    `json:"location"`
	IsPublic       bool      `json:"is_public"`
	CreatorID      uint      `json:"creator_id"`
	OrganizationID uint      `json:"organization_id"`
	Date           time.Time `json:"date"`
	StartTime      string    `json:"start_time"`
	EndTime        string    `json:"end_time"`
}
