package main

import (
	"fmt"

	"github.com/TeeHaoBin/TicketLah/backend/config"
	"github.com/TeeHaoBin/TicketLah/backend/db"
	"github.com/TeeHaoBin/TicketLah/backend/handlers"
	"github.com/TeeHaoBin/TicketLah/backend/middlewares"
	"github.com/TeeHaoBin/TicketLah/backend/repositories"
	"github.com/TeeHaoBin/TicketLah/backend/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
    envConfig := config.NewEnvConfig()
    db := db.Init(envConfig, db.DBMigrator)

    app := fiber.New(fiber.Config{
        AppName: "TicketLah Backend",
        ServerHeader: "Fiber: TicketLah-Backend-Server",
    })

    // Repository
    eventRepository := repositories.NewEventRepository(db)
    ticketRepository := repositories.NewTicketRepository(db)
    authRepository := repositories.NewAuthRepository(db)

    // Service
    authService := services.NewAuthService(authRepository)

    // Routing
    server := app.Group("/api")
    handlers.NewAuthHandler(server.Group("/auth"), authService)

    privateRoutes := server.Use(middlewares.AuthProtected(db))

    // Handlers
    handlers.NewEventRepository(privateRoutes.Group("/event"), eventRepository)
    handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

    app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
