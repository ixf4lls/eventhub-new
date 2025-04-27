package middleware

import (
	"eventhub-backend/internal/service"
	customJwt "eventhub-backend/pkg/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
	jwtManager customJwt.Manager
}

func NewMiddleware(jwtManager customJwt.Manager) *Middleware {
	return &Middleware{jwtManager: jwtManager}
}

func (m *Middleware) AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Токен невалиден или отстутствует")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := m.jwtManager.ParseToken(tokenStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("userID", userID)

		return next(c)
	}
}

func (m *Middleware) IsOrganizationCreator(orgService *service.OrganizationService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Get("userID").(uint)
			orgIDstr := c.Param("id")
			orgID, err := strconv.ParseUint(orgIDstr, 10, 64)

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Некорректный ID организации")
			}

			creatorID, err := orgService.GetCreator(uint(orgID))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при проверке прав доступа")
			}

			if userID != creatorID {
				return echo.NewHTTPError(http.StatusForbidden, "У вас нет прав для управления этой организацией")
			}

			return next(c)
		}
	}
}
