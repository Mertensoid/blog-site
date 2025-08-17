package home

import (
	"blog-site/internal/register"
	"blog-site/package/bcrypt"
	"blog-site/package/templadapter"
	"blog-site/views"
	"blog-site/views/pages"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *register.UsersRepository
	cryptograf *bcrypt.Crypto
}

func NewHandler(router fiber.Router,
	logger *zerolog.Logger,
	repository *register.UsersRepository,
	cryptograf *bcrypt.Crypto) {
	h := &HomeHandler{
		router:     router,
		logger:     logger,
		repository: repository,
		cryptograf: cryptograf,
	}
	h.router.Get("/", h.home)
	h.router.Get("/register", h.register)
	h.router.Get("/error", h.error)
	h.router.Get("/test", h.test)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return templadapter.Render(c, component)
}

func (h *HomeHandler) register(c *fiber.Ctx) error {
	component := pages.Register()
	return templadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().Msg("Hello")
	return c.SendString("Error")
}

func (h *HomeHandler) test(c *fiber.Ctx) error {
	h.logger.Info().Msg("UserTest")
	user := h.repository.CheckUser("mma@mma.mma", "12345")

	return c.SendString(fmt.Sprintf("User %d %s %s %s %s", user.Id, user.Email, user.Password, user.Name, user.RegTime))
}
