package registration

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/rs/zerolog"

	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/pkg/validator"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type RegisterHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *UserRepository
	store        *session.Store
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *UserRepository, store *session.Store) {
	h := RegisterHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	api := h.router.Group("/api")
	api.Post("/signUp", h.signUp)
	api.Post("/login", h.login)
	api.Get("/logout", h.logout)
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

	userName, err := h.repository.loginUser(form.Email, form.Password)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component := components.Notification(err.Error(), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	session, err := h.store.Get(c)
	if err != nil {
		log.Printf("Ошибка получения сессии: %v", err)
		return c.Redirect("/login", fiber.StatusSeeOther)
	}
	session.Set("userName", userName)
	if err := session.Save(); err != nil {
		log.Printf("Ошибка получения сессии: %v", err)
		return c.Redirect("/login", fiber.StatusSeeOther)
	}
	c.Response().Header.Add("HX-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *RegisterHandler) logout(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		log.Printf("Ошибка получения сессии: %v", err)
		return c.Redirect("/login", fiber.StatusSeeOther)
	}
	if err := session.Destroy(); err != nil {
		log.Printf("Ошибка удаления сессии: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при выходе из аккаунта")
	}
	c.Response().Header.Add("HX-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}
