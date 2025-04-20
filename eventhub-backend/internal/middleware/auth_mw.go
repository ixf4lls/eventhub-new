package middleware

import (
	customJwt "eventhub-backend/pkg/jwt"
	"net/http"
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
