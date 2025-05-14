package handlers

import (
	"errors"
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/repository"
	"eventhub-backend/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventHandler struct {
	eventService        *service.EventService
	userService         *service.UserService
	notificationService *service.NotificationService
}

func NewEventHandler(eventService *service.EventService, userService *service.UserService, notificationService *service.NotificationService) *EventHandler {
	return &EventHandler{eventService: eventService, userService: userService, notificationService: notificationService}
}

func (h *EventHandler) GetAllUser(c echo.Context) error {
	userID := c.Get("userID").(uint)

	joinedEvents, openEvents, organizationsEvents, err := h.eventService.GetAllUser(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении событий")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"joined_events":        joinedEvents,
		"open_events":          openEvents,
		"organizations_events": organizationsEvents,
	})
}

func (h *EventHandler) Join(c echo.Context) error {
	userID := c.Get("userID").(uint)

	eventIDStr := c.Param("id")
	if eventIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.eventService.Join(userID, uint(eventID)); err != nil {
		if err.Error() == "event not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "Мероприятие не существует")
		}

		if err.Error() == "user already joined" {
			return echo.NewHTTPError(http.StatusBadRequest, "Пользователь уже записан на это мероприятие")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при присоединении к мероприятию")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EventHandler) Quit(c echo.Context) error {
	userID := c.Get("userID").(uint)

	eventIDStr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.eventService.Quit(userID, uint(eventID)); err != nil {
		if err.Error() == "event not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "Мероприятие не существует")
		}

		if err.Error() == "not joined" {
			return echo.NewHTTPError(http.StatusBadRequest, "Пользователь не записан на это мероприятие")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при отмене записи на мероприятие")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EventHandler) Create(c echo.Context) error {
	userID := c.Get("userID").(uint)

	orgIDStr := c.Param("id")
	orgID, err := strconv.ParseUint(orgIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	var input domain.CreateEventInput
	if err := c.Bind(&input); err != nil {
		log.Println(input)
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.eventService.Create(input, userID, uint(orgID)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при создании мероприятия")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EventHandler) GetByID(c echo.Context) error {
	userID := c.Get("userID").(uint)
	eventIDstr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDstr, 10, 64)
	if err != nil || eventIDstr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	event, isCreator, err := h.eventService.GetByID(uint(eventID), userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Мероприятие не найдено")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении данных мероприятия")
	}

	isJoined, err := h.eventService.IsUserJoined(userID, uint(eventID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при проверке участия в мероприятии")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"event":      event,
		"is_creator": isCreator,
		"is_joined":  isJoined,
	})
}

func (h *EventHandler) Delete(c echo.Context) error {
	userID := c.Get("userID").(uint)
	eventIDstr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDstr, 10, 64)
	if err != nil || eventIDstr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	participants, err := h.eventService.GetParticipantIDs(uint(eventID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при обновлении мероприятия")
	}

	for _, participant := range participants {
		if err := h.notificationService.Create(participant, uint(eventID), "cancel"); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при обновлении мероприятия")
		}
	}

	if err := h.eventService.Delete(userID, uint(eventID)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при удалении мероприятия")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EventHandler) Update(c echo.Context) error {
	userID := c.Get("userID").(uint)

	orgIDStr := c.Param("id")
	orgID, err := strconv.ParseUint(orgIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	eventIDstr := c.Param("event_id")
	eventID, err := strconv.ParseUint(eventIDstr, 10, 64)
	if err != nil || eventIDstr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный ID мероприятия")
	}

	var input domain.CreateEventInput
	if err := c.Bind(&input); err != nil {
		log.Println(input)
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.eventService.Update(userID, uint(eventID), uint(orgID), input); err != nil {
		if err.Error() == "access denied" {
			return echo.NewHTTPError(http.StatusForbidden, "У вас нет прав для обновления этого мероприятия")
		}
		if err.Error() == "event not exists" {
			return echo.NewHTTPError(http.StatusNotFound, "Мероприятие не найдено")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при обновлении мероприятия")
	}

	participants, err := h.eventService.GetParticipantIDs(uint(eventID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при обновлении мероприятия")
	}

	for _, participant := range participants {
		if err := h.notificationService.Create(participant, uint(eventID), "reschedule"); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при обновлении мероприятия")
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Мероприятие успешно обновлено"})
}

func (h *EventHandler) UpdateSearchIndex(c echo.Context) error {
	err := h.eventService.UpdateSearchIndex()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при индексации")
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Индексация успешно завершена"})
}

func (h *EventHandler) GetParticipants(c echo.Context) error {
	eventIDstr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDstr, 10, 64)
	if err != nil || eventIDstr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный ID мероприятия")
	}

	participantIDs, err := h.eventService.GetParticipantIDs(uint(eventID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении участников мероприятия")
	}

	var users []repository.UserResponse
	for _, participantID := range participantIDs {
		user, err := h.userService.GetByID(participantID)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"participants": users})
}
