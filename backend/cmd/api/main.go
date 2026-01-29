package main

import (
	"fmt"

	"github.com/TeeHaoBin/TicketLah/backend/config"
	"github.com/TeeHaoBin/TicketLah/backend/db"
	"github.com/TeeHaoBin/TicketLah/backend/handlers"
	"github.com/TeeHaoBin/TicketLah/backend/repositories"
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

    // Routing
    server := app.Group("/api")

    // Handlers
    handlers.NewEventRepository(server.Group("/event"), eventRepository)
    handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

    app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
