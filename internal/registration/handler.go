package registration

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/rs/zerolog"

	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/pkg/validator"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *UserRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *UserRepository) {
	h := RegisterHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	api := h.router.Group("/api")
	api.Post("/signUp", h.signUp)
	api.Post("/login", h.login)
}

func (h *RegisterHandler) signUp(c *fiber.Ctx) error {
	form := RegistrationForm{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Name", Field: form.Name, Message: "Имя не задано"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password, Message: "Пароль не задан"})
	var component templ.Component
	if len(errors.Errors) > 0 {
		h.customLogger.Error().Msg(validator.FormatErrors(errors))
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	err := h.repository.addUser(form)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification(err.Error(), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	component = components.Notification("✅ Регистрация прошла успешно!", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *RegisterHandler) login(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Введите email"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password, Message: "Введите пароль"},
	)
	if len(errors.Errors) > 0 {
		h.customLogger.Error().Msg(validator.FormatErrors(errors))
		component := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}

	username, err := h.repository.loginUser(form.Email, form.Password)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component := components.Notification(err.Error(), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}

	component := components.Notification(fmt.Sprintf("Добро пожаловать, %s!", username), components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}
