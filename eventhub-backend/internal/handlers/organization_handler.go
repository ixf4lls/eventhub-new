package handlers

import (
	"eventhub-backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrganizationHandler struct {
	organizationService *service.OrganizationService
}

type createOrgRequest struct {
	Name string `json:"organization_name"`
}

func NewOrganizationHandler(organizationService *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{organizationService: organizationService}
}

func (h *OrganizationHandler) Create(c echo.Context) error {
	userID := c.Get("userID").(uint)
	var req createOrgRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Невалидный JSON")
	}

	name := req.Name
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.organizationService.Create(name, userID); err != nil {
		if err.Error() == "organization_already_exists" {
			return echo.NewHTTPError(http.StatusConflict, "Организация с таким именем уже существует")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при создании организации")
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *OrganizationHandler) GetAll(c echo.Context) error {
	userID := c.Get("userID").(uint)

	joined, founded, err := h.organizationService.GetAll(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении организаций пользователя")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"joined":  joined,
		"founded": founded,
	})
}

func (h *OrganizationHandler) JoinByCode(c echo.Context) error {
	userID := c.Get("userID").(uint)

	code := c.Param("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	if err := h.organizationService.JoinByCode(userID, code); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Ошибка при присоединении к организации")
	}

	return c.NoContent(http.StatusNoContent)
}
