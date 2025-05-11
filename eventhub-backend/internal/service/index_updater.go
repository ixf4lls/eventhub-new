package service

import (
	"log"
	"time"
)

func (s *EventService) StartIndexUpdater() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			err := s.UpdateSearchIndex()
			if err != nil {
				log.Println("Ошибка при обновлении индексов:", err)
			}
		}
	}()
}
