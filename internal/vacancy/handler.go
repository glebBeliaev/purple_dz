package vacancy

import (
	"github.com/a-h/templ"
	"github.com/glebbeliaev/purple_dz/pkg/tadaptor"
	"github.com/glebbeliaev/purple_dz/pkg/validator"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router        fiber.Router
	cusstomLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, cusstomLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:        router,
		cusstomLogger: cusstomLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateFrom{
		Email:    c.FormValue("email"),
		Role:     c.FormValue("role"),
		Name:     c.FormValue("name"),
		Type:     c.FormValue("type"),
		Salary:   c.FormValue("salary"),
		Location: c.FormValue("location"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Должность не задана"},
		&validators.StringIsPresent{Name: "Name", Field: form.Name, Message: "Название компании не задано"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "Сфера компании не задана"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Заработная плата не задана"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Место работы не задано"})
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadaptor.Render(c, component)
	}
	h.cusstomLogger.Info().Msg(form.Email)
	component = components.Notification("Vacancy created!", components.NotificationSuccess)
	return tadaptor.Render(c, component)
}
