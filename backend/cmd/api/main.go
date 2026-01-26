package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/TeeHaoBin/TicketLah/backend/handlers"
    "github.com/TeeHaoBin/TicketLah/backend/repositories"
)

func main() {
    app := fiber.New(fiber.Config{
        AppName: "TicketLah Backend",
        ServerHeader: "Fiber: TicketLah-Backend-Server",
    })

    // Repository
    eventRepository := repositories.NewEventRepository(nil)

    // Routing
    server := app.Group("/api")

    // Handlers
    handlers.NewEventRepository(server.Group("/event"), eventRepository)

    app.Listen(":3000")
}
