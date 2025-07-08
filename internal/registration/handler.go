package registration

import (
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
	customLogger zerolog.Logger
	repository   *UserRepository
}

func NewHandler(router fiber.Router) {
	h := RegisterHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Post("/signUp", h.signUp)
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
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("✅ Регистрация прошла успешно!", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
