package news

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/pkg/validator"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type CreateNewsHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *NewsRepository
	store        *session.Store
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *NewsRepository, store *session.Store) {
	h := CreateNewsHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	api := h.router.Group("/news")
	api.Post("/create", h.createNews)
}

func (h *CreateNewsHandler) createNews(c *fiber.Ctx) error {

	Author := c.Locals("userName").(string)

	form := NewsForm{
		Title:    c.FormValue("title"),
		Preview:  c.FormValue("preview"),
		Content:  c.FormValue("content"),
		Cover:    c.FormValue("cover"),
		Category: c.FormValue("category"),
		Author:   Author,
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "title", Field: form.Title, Message: "Заголовок не задан"},
		&validators.StringIsPresent{Name: "preview", Field: form.Preview, Message: "Превью не задано"},
		&validators.StringIsPresent{Name: "content", Field: form.Content, Message: "Текст не задан"},
		&validators.StringIsPresent{Name: "cover", Field: form.Cover, Message: "Картинка не задана"})
	var component templ.Component

	if len(errors.Errors) > 0 {
		h.customLogger.Error().Msg(validator.FormatErrors(errors))
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}

	err := h.repository.createNews(form)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification(err.Error(), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	component = components.Notification("✅ Новость создана успешно!", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}
