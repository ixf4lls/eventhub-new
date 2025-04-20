package main

import (
	"eventhub-backend/internal/config"
	"eventhub-backend/internal/database"
	"eventhub-backend/internal/handlers"
	"eventhub-backend/internal/middleware"
	"eventhub-backend/internal/repository"
	"eventhub-backend/internal/service"
	customJwt "eventhub-backend/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.ConnectToDB()
	database.MigrageDB(db)

	e := echo.New()

	cfg := config.Load()
	var jwtManager customJwt.Manager = customJwt.NewJwtManager(cfg.JwtSecretKey)

	// repos
	userRepo := repository.NewGormUserRepository(db)
	eventRepo := repository.NewGormEventRepository(db)
	organizationRepo := repository.NewGormOrganizationRepository(db)

	// services
	authService := service.NewAuthService(cfg, userRepo, jwtManager)
	registerService := service.NewRegisterService(userRepo)
	eventService := service.NewEventService(*eventRepo)
	organizationService := service.NewOrganizationService(*organizationRepo)

	// handlers
	authHandler := handlers.NewAuthHandler(authService, registerService)
	eventHandler := handlers.NewEventHandler(eventService)
	organizationHandler := handlers.NewOrganizationHandler(organizationService)

	// middleware
	authMW := middleware.NewMiddleware(jwtManager)

	api := e.Group("/api")

	// public
	api.POST("/refresh", authHandler.Refresh)   // POST /api/refresh
	api.POST("/register", authHandler.Register) // POST /api/register
	api.POST("/login", authHandler.Login)       // POST /api/login

	// authorized
	auth := api.Group("", authMW.AuthRequired)

	events := auth.Group("/events")
	events.GET("", eventHandler.GetAll)                 // GET    /api/events
	events.POST("/:event_id/join", eventHandler.Join)   // POST   /api/events/:id/join
	events.DELETE("/:event_id/quit", eventHandler.Quit) // DELETE /api/events/:id/quit
	events.POST("/create", eventHandler.Create)         // POST /api/events/create

	organizations := auth.Group("/organizations")
	organizations.GET("", organizationHandler.GetAll)                 // GET  /api/organizations
	organizations.POST("/create", organizationHandler.Create)         // POST /api/organizations/create
	organizations.POST("/join/:code", organizationHandler.JoinByCode) // POST /api/organizations/join/:code

	e.Start(":3000")
}
