package registration

import (
	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"

	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/pkg/validator"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := RegisterHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Post("/signUp", h.signUp)
}

func (h *RegisterHandler) signUp(c *fiber.Ctx) error {
	form := RegistrationFrom{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Name", Field: form.Name, Message: "Название компании не задано"},
		&validators.StringIsPresent{Name: "Type", Field: form.Password, Message: "Сфера компании не задана"})
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Vacancy created!", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
