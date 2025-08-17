package register

import (
	"blog-site/package/bcrypt"
	"blog-site/package/templadapter"
	"blog-site/package/validator"
	"blog-site/views/components"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type RegisterHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *UsersRepository
	cryptograf *bcrypt.Crypto
}

func NewHandler(router fiber.Router,
	logger *zerolog.Logger,
	repository *UsersRepository,
	cryptograf *bcrypt.Crypto) {
	h := &RegisterHandler{
		router:     router,
		logger:     logger,
		repository: repository,
		cryptograf: cryptograf,
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

	if len(errors.Errors) > 0 {
		component := components.Notification(validator.FormatErrors(*errors), components.NotificationFail)
		return templadapter.Render(c, component)
	}
	err := h.repository.addUser(form)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		component := components.Notification("Ошибка на сервере при попытке регистрации", components.NotificationFail)
		return templadapter.Render(c, component)
	}
	component := components.Notification("Регистрация успешно выполнена", components.NotificationSuccess)
	return templadapter.Render(c, component)
}
