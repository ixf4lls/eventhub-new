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
	events.GET("", eventHandler.GetAll)           // GET    /api/events
	events.GET("/:id", eventHandler.GetByID)      // GET /api/events/:id
	events.POST("/:id/join", eventHandler.Join)   // POST   /api/events/:id/join
	events.DELETE("/:id/quit", eventHandler.Quit) // DELETE /api/events/:id/quit

	organizations := auth.Group("/organizations")
	organizations.GET("", organizationHandler.GetAll)                 // GET  /api/organizations
	organizations.GET("/:id/events", organizationHandler.GetEvents)   // GET /api/organizations/:id/events
	organizations.GET("/:id", organizationHandler.GetByID)            // GET /api/organizations/:id
	organizations.GET("/:id/members", organizationHandler.GetMembers) // GET /api/organizations/:id/members
	organizations.POST("", organizationHandler.Create)                // POST /api/organizations/
	organizations.POST("/join/:code", organizationHandler.JoinByCode) // POST /api/organizations/join/:code

	orgCreator := organizations.Group("/:id", authMW.IsOrganizationCreator(organizationService))
	orgCreator.POST("/events", eventHandler.Create) // POST /api/organizations/:id/events

	e.Start(":3000")
}
