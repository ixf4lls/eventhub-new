package service

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

const (
	reminder1dBefore = time.Hour * 23
	reminder1hBefore = time.Hour
)

func (s *NotificationService) StartScheduler() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			err := s.checkAndSendReminders()
			if err != nil {
				log.Println("Ошибка при отправке напоминаний:", err)
			}
		}
	}()
}

func (s *NotificationService) checkAndSendReminders() error {
	events, err := s.eventRepo.GetUpcoming()
	if err != nil {
		return err
	}

	now := time.Now()

	for _, event := range events {
		status := event.Status
		if status == "deleted" || status == "completed" {
			continue
		}

		eventStartTime, err := time.Parse("15:04:05", event.StartTime)
		if err != nil {
			log.Printf("Ошибка при парсинге времени события %s: %v", event.StartTime, err)
			continue
		}

		eventTime := combineDateTime(event.Date, eventStartTime)
		duration := eventTime.Sub(now)

		var reminderType string
		switch {
		case duration >= reminder1dBefore && duration < time.Hour*24:
			reminderType = "reminder_1d"
		case duration >= reminder1hBefore && duration < time.Hour*2:
			reminderType = "reminder_1h"
		default:
			continue
		}

		participants, err := s.eventRepo.GetParticipants(event.ID)
		if err != nil {
			log.Printf("Ошибка при получении участников для события %d: %v", event.ID, err)
			continue
		}

		for _, userID := range participants {
			s.trySendReminder(event.ID, userID, reminderType, now)
		}
	}

	return nil
}

func (s *NotificationService) trySendReminder(eventID, userID uint, reminderType string, now time.Time) {
	alreadySent, err := s.notificationRepo.Exists(eventID, userID, reminderType)
	if (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || alreadySent {
		return
	}

	err = s.Create(userID, eventID, reminderType)
	log.Printf("Sending %s reminder to user %d for event %d", reminderType, userID, eventID)
	if err != nil {
		log.Printf("Ошибка при создании %s уведомления: %v", reminderType, err)
	}
}

func combineDateTime(date time.Time, startTime time.Time) time.Time {
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		startTime.Hour(), startTime.Minute(), startTime.Second(), 0, date.Location(),
	)
}
