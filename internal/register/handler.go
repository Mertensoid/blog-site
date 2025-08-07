package register

import (
	"blog-site/package/templadapter"
	"blog-site/package/validator"
	"blog-site/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type RegisterHandler struct {
	router fiber.Router
	logger zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	h := &RegisterHandler{
		router: router,
		logger: *logger,
	}
	h.router.Post("/api/register", h.register)
}

func (h *RegisterHandler) register(c *fiber.Ctx) error {
	form := RegisterData{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{
			Name:    "Name",
			Field:   form.Name,
			Message: "Имя не задано",
		},
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан",
		},
		&validators.StringIsPresent{
			Name:    "Password",
			Field:   form.Password,
			Message: "Пароль не задан",
		},
	)

	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(*errors), components.NotificationFail)
	} else {
		component = components.Notification("Успешная регистрация", components.NotificationSuccess)
	}
	return templadapter.Render(c, component)
}
