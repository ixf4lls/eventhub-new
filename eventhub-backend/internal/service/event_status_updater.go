package service

import (
	"log"
	"time"
)

func (s *NotificationService) StartEventStatusUpdater() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			eventsIDsToMarkCompleted, err := s.eventRepo.GetEventIDsToMarkCompleted()
			if err != nil {
				log.Println("Ошибка при получении мероприятий для обновления статуса")
			}

			err = s.eventRepo.MarkEventsCompleted(eventsIDsToMarkCompleted)
			if err != nil {
				log.Println("Ошибка при обновлении статусов мероприятий:", err)
			}

		}
	}()
}
