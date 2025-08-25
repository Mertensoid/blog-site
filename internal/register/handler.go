package register

import (
	"blog-site/package/bcrypt"
	"blog-site/package/templadapter"
	"blog-site/package/validator"
	"blog-site/views/components"
	"net/http"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type RegisterHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *UsersRepository
	cryptograf *bcrypt.Crypto
	store      *session.Store
}

func NewHandler(router fiber.Router,
	logger *zerolog.Logger,
	repository *UsersRepository,
	cryptograf *bcrypt.Crypto,
	store *session.Store) {
	h := &RegisterHandler{
		router:     router,
		logger:     logger,
		repository: repository,
		cryptograf: cryptograf,
		store:      store,
	}
	h.router.Post("/api/register", h.register)
	h.router.Post("/login", h.checkUser)
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
		return templadapter.Render(c, component, http.StatusBadRequest)
	}
	err := h.repository.addUser(form)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		component := components.Notification("Ошибка на сервере при попытке регистрации", components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}
	component := components.Notification("Регистрация успешно выполнена", components.NotificationSuccess)
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *RegisterHandler) checkUser(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не введен или не верный",
		},
		&validators.StringIsPresent{
			Name:    "Password",
			Field:   form.Password,
			Message: "Пароль не введен",
		},
	)

	if len(errors.Errors) > 0 {
		component := components.Notification(validator.FormatErrors(*errors), components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}
	user, err := h.repository.checkUser(form)
	if err != nil {
		if err.Error() == "Incorrect password" {
			h.logger.Error().Msg(err.Error())
			component := components.Notification("Неверный пароль", components.NotificationFail)
			return templadapter.Render(c, component, http.StatusBadRequest)
		}
		h.logger.Error().Msg(err.Error())
		component := components.Notification("Ошибка на сервере", components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}

	// Создать сессию
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	session.Set("email", user.Email)
	session.Set("name", user.Name)
	if err := session.Save(); err != nil {
		panic(err)
	}

	// Перенаправление на главную
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}
