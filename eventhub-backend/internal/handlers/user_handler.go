package handlers

import (
	"eventhub-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserData(c echo.Context) error {
	userID := c.Get("userID").(uint)
	userResponse, err := h.userService.GetByID(uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении информации о пользователе")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"user": userResponse})
}

func (h *UserHandler) GetByID(c echo.Context) error {
	userIDstr := c.Param("id")
	if userIDstr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	userResponse, err := h.userService.GetByID(uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении информации о пользователе")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user": userResponse})
}
