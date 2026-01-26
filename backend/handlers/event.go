package handlers

import (
	"github.com/TeeHaoBin/TicketLah/backend/models"
	"github.com/gofiber/fiber/v2"
	"context"
	"time"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map {
			"status": "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map {
		"status": "success",
		"message": "Get events successfully",
		"data": events,
	})
}

func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}

func NewEventRepository(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
}