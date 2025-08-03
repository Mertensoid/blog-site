package home

import (
	"blog-site/package/templadapter"
	"blog-site/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	logger zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	h := &HomeHandler{
		router: router,
		logger: *logger,
	}
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return templadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().Msg("Hello")
	return c.SendString("Error")
}
