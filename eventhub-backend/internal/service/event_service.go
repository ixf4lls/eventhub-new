package service

import (
	"errors"
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/repository"
	"time"
)

type EventService struct {
	eventRepo repository.GormEventRepository
}

func NewEventService(eventRepo repository.GormEventRepository) *EventService {
	return &EventService{eventRepo: eventRepo}
}

func (s *EventService) GetAll(userID uint) ([]repository.EventResponse, []repository.EventResponse, error) {
	return s.eventRepo.GetAll(userID)
}

func (s *EventService) Join(userID, eventID uint) error {
	exists, err := s.eventRepo.IsEventExist(eventID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("event not found")
	}

	joined, err := s.eventRepo.IsUserJoined(userID, eventID)
	if err != nil {
		return err
	}
	if joined {
		return errors.New("user already joined")
	}

	return s.eventRepo.Join(userID, eventID)
}

func (s *EventService) Quit(userID, eventID uint) error {
	exists, err := s.eventRepo.IsEventExist(eventID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("event not found")
	}

	joined, err := s.eventRepo.IsUserJoined(userID, eventID)
	if err != nil {
		return err
	}
	if !joined {
		return errors.New("user not joined")
	}

	return s.eventRepo.Quit(userID, eventID)
}

func (s *EventService) Create(input domain.CreateEventInput, founderID uint) error {
	today := time.Now().Truncate(24 * time.Hour)
	if input.Date.Before(today) {
		return errors.New("date in past")
	}

	if input.StartTime.After(input.EndTime) {
		return errors.New("incorrect time")
	}

	return s.eventRepo.Create(input, founderID)
}
