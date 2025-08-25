package home

import (
	"blog-site/internal/register"
	"blog-site/package/bcrypt"
	"blog-site/package/templadapter"
	"blog-site/views"
	"blog-site/views/pages"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *register.UsersRepository
	cryptograf *bcrypt.Crypto
	store      *session.Store
}

func NewHandler(router fiber.Router,
	logger *zerolog.Logger,
	repository *register.UsersRepository,
	cryptograf *bcrypt.Crypto,
	store *session.Store) {
	h := &HomeHandler{
		router:     router,
		logger:     logger,
		repository: repository,
		cryptograf: cryptograf,
		store:      store,
	}
	h.router.Get("/", h.home)
	h.router.Get("/entrance", h.login)
	h.router.Get("/register", h.register)
	h.router.Get("/error", h.error)
	h.router.Get("/logout", h.logout)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := pages.Login()
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) logout(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	session.Delete("email")
	session.Delete("name")
	if err := session.Save(); err != nil {
		panic(err)
	}
	return c.Redirect("/")
}

func (h *HomeHandler) register(c *fiber.Ctx) error {
	component := pages.Register()
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().Msg("Hello")
	return c.SendString("Error")
}
