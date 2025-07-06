package vacancy

import (
	"net/http"

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
	repository    *VacancyRepository
}

func NewHandler(router fiber.Router, cusstomLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &VacancyHandler{
		router:        router,
		cusstomLogger: cusstomLogger,
		repository:    repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateFrom{
		Email:       c.FormValue("email"),
		Role:        c.FormValue("role"),
		Company:     c.FormValue("company"),
		CompanyType: c.FormValue("type"),
		Salary:      c.FormValue("salary"),
		Location:    c.FormValue("location"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Должность не задана"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Название компании не задано"},
		&validators.StringIsPresent{Name: "Type", Field: form.CompanyType, Message: "Сфера компании не задана"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Заработная плата не задана"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Место работы не задано"})
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadaptor.Render(c, component, http.StatusBadRequest)
	}
	err := h.repository.addVacancy(form)
	if err != nil {
		h.cusstomLogger.Error().Msg(err.Error())
		component = components.Notification("Произошла ошибка на сервере, попробуйте позднее", components.NotificationFail)
		return tadaptor.Render(c, component, http.StatusBadRequest)
	}
	h.cusstomLogger.Info().Msg(form.Email)
	component = components.Notification("Vacancy created!", components.NotificationSuccess)
	return tadaptor.Render(c, component, http.StatusOK)
}
