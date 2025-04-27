package handlers

import (
	"eventhub-backend/internal/service"
	"net/http"
	"strconv"

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
		if err.Error() == "organization not found" {
			return echo.NewHTTPError(http.StatusNotFound, "Код прилашения не существует")
		}

		if err.Error() == "user already a member of the organization" {
			return echo.NewHTTPError(http.StatusConflict, "Вы уже присоединились к этой организации")
		}

		if err.Error() == "user is a creator of this organization" {
			return echo.NewHTTPError(http.StatusConflict, "Вы не можете присоединиться к своей же организации")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при присоединении к организации")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *OrganizationHandler) GetEvents(c echo.Context) error {
	orgIDstr := c.Param("id")
	orgID, err := strconv.ParseUint(orgIDstr, 10, 16)
	if orgIDstr == "" || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	activeEvents, completedEvents, err := h.organizationService.GetEvents(uint(orgID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении мероприятий организации")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"active":    activeEvents,
		"completed": completedEvents,
	})
}

func (h *OrganizationHandler) GetByID(c echo.Context) error {
	userID := c.Get("userID").(uint)

	orgIDstr := c.Param("id")
	orgID, err := strconv.ParseUint(orgIDstr, 10, 16)
	if orgIDstr == "" || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	organization, isCreator, err := h.organizationService.GetByID(uint(orgID), userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении информации об организации")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"organization": organization,
		"is_creator":   isCreator,
	})
}

func (h *OrganizationHandler) GetMembers(c echo.Context) error {
	orgIDstr := c.Param("id")
	orgID, err := strconv.ParseUint(orgIDstr, 10, 16)
	if orgIDstr == "" || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный запрос")
	}

	members, err := h.organizationService.GetMembers(uint(orgID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении списка участников")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"members": members,
	})
}
