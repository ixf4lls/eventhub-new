package handlers

import (
	"eventhub-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationService *service.NotificationService
}

func NewNotificationHandler(notificationService *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

func (h *NotificationHandler) Create(c echo.Context) error {
	userID := c.Get("userID").(uint)

	eventIDStr := c.Param("id")
	if eventIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}
	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	eventType := c.Param("type")
	if eventType == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.notificationService.Create(userID, uint(eventID), eventType); err != nil {
		if err.Error() == "incorrect msg type" {
			return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
		}
		if err.Error() == "event do not exists" {
			return echo.NewHTTPError(http.StatusBadRequest, "Мероприятие не существует")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при создании уведомления")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *NotificationHandler) GetAll(c echo.Context) error {
	userID := c.Get("userID").(uint)

	notifications, err := h.notificationService.GetAll(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении уведомлений")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"notifications": notifications,
	})
}
