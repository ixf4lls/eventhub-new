package service

import (
	"errors"
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/repository"
	"fmt"
	"time"
)

type NotificationService struct {
	notificationRepo repository.GormNotificationRepository
	eventRepo        repository.GormEventRepository
}

func NewNotificationService(notificationRepo repository.GormNotificationRepository, eventRepo repository.GormEventRepository) *NotificationService {
	return &NotificationService{notificationRepo: notificationRepo, eventRepo: eventRepo}
}

func (s *NotificationService) Create(userID, eventID uint, msgType string) error {
	if msgType != "reminder_1d" && msgType != "reminder_1h" && msgType != "cancel" && msgType != "reschedule" {
		return errors.New("incorrect msg type")
	}

	exists, err := s.eventRepo.IsEventExist(eventID)
	if err != nil {
		return nil
	}

	if !exists {
		return errors.New("event do not exists")
	}

	return s.notificationRepo.Create(userID, eventID, msgType)
}

func (s *NotificationService) GetAll(userID uint) ([]domain.Notification, error) {
	notifications, err := s.notificationRepo.GetAll(userID)

	if err != nil {
		return nil, err
	}

	var notificationsResponse []domain.Notification
	for _, ntf := range notifications {
		event, _, err := s.eventRepo.GetByID(ntf.EventID, ntf.UserID)
		if err != nil {
			continue
		}

		parsedDate, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			continue
		}
		parsedTime, err := time.Parse("15:04:05", event.StartTime)
		if err != nil {
			continue
		}
		eventDateTime := time.Date(
			parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
			parsedTime.Hour(), parsedTime.Minute(), 0, 0, time.Local,
		)
		months := []string{
			"января", "февраля", "марта", "апреля", "мая", "июня",
			"июля", "августа", "сентября", "октября", "ноября", "декабря",
		}
		formatted := fmt.Sprintf("%d %s, %02d:%02d",
			eventDateTime.Day(),
			months[eventDateTime.Month()-1],
			eventDateTime.Hour(),
			eventDateTime.Minute(),
		)

		var message, info string
		switch ntf.Type {
		case "reminder_1d":
			message = fmt.Sprintf("🥳 «%s» уже завтра!", event.Title)
			info = "До встречи на мероприятии! Готовься, будет классно!"
		case "reminder_1h":
			message = fmt.Sprintf("🥳 «%s» уже через час!", event.Title)
			info = "До встречи на мероприятии! Готовься, будет классно!"
		case "cancel":
			message = fmt.Sprintf("😔 Мероприятие «%s» отменено", event.Title)
			info = "К сожалению, мероприятие не состоится. Надеемся увидеть тебя на других событиях!"
		case "reschedule":
			message = fmt.Sprintf("❗ Мероприятие «%s» перенесено", event.Title)
			info = fmt.Sprintf("Новое время: %s. Мы ждем тебя!", formatted)
		default:
			continue
		}

		notificationsResponse = append(notificationsResponse, domain.Notification{
			Message: message,
			Info:    info,
		})
	}

	return notificationsResponse, nil
}
