package handlers

import (
	"eventhub-backend/internal/domain"
	"eventhub-backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthHandler struct {
	authService     *service.AuthService
	registerService *service.RegisterService
}

func NewAuthHandler(authService *service.AuthService, registerService *service.RegisterService) *AuthHandler {
	return &AuthHandler{
		authService:     authService,
		registerService: registerService,
	}
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	var req RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	access, refresh, err := h.authService.RefreshTokens(req.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *AuthHandler) Register(c echo.Context) error {
	var input domain.RegisterInput

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.registerService.Register(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Пользователь успешно зарегистрирован",
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var input domain.LoginInput

	err := c.Bind(&input)
	if err != nil || input.Username == "" || input.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	accessToken, refreshToken, err := h.authService.Login(input)
	if err != nil {
		if err.Error() == "invalid credentials" {
			return echo.NewHTTPError(http.StatusBadRequest, "Неправильный логин или пароль")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при попытке входа")
	}

	return c.JSON(http.StatusOK, map[string]string{"access_token": accessToken, "refresh_token": refreshToken})
}
