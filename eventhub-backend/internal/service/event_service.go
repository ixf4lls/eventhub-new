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

func (s *EventService) Create(input domain.CreateEventInput, creatorID, orgID uint) error {
	today := time.Now().Truncate(24 * time.Hour)
	if input.Date.Before(today) {
		return errors.New("date in past")
	}

	inputStartTime, _ := time.Parse("15:04:05", input.StartTime)

	inputEndTime, _ := time.Parse("15:04:05", input.EndTime)

	if inputStartTime.After(inputEndTime) {
		return errors.New("incorrect time")
	}

	return s.eventRepo.Create(input, creatorID, orgID)
}

func (s *EventService) GetByID(eventID, userID uint) (repository.EventResponse, bool, error) {
	return s.eventRepo.GetByID(eventID, userID)
}

func (s *EventService) IsUserJoined(userID, eventID uint) (bool, error) {
	return s.eventRepo.IsUserJoined(userID, eventID)
}

func (s *EventService) IsUserCreator(userID, eventID uint) (bool, error) {
	return s.eventRepo.IsUserCreator(userID, eventID)
}

func (s *EventService) Delete(userID, eventID uint) error {
	isCreator, err := s.eventRepo.IsUserCreator(userID, eventID)
	if err != nil {
		return err
	}

	if !isCreator {
		return errors.New("access denied")
	}

	return s.eventRepo.Delete(eventID)
}
