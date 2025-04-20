package handlers

import (
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventService *service.EventService
}

func NewEventHandler(eventService *service.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) GetAll(c echo.Context) error {
	userID := c.Get("userID").(uint)

	joinedEvents, openEvents, err := h.eventService.GetAll(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении событий")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"joined_events": joinedEvents,
		"open_events":   openEvents,
	})
}

func (h *EventHandler) Join(c echo.Context) error {
	userID := c.Get("userID").(uint)

	eventIDStr := c.Param("event_id")
	if eventIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный event_id")
	}

	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный event_id")
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

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный event_id")
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

	var input domain.CreateEventInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.eventService.Create(input, userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при создании мероприятия")
	}

	return c.NoContent(http.StatusNoContent)
}
