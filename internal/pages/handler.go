package pages

import (
	"math"
	"net/http"

	"github.com/glebbeliaev/purple_dz/internal/vacancy"
	"github.com/glebbeliaev/purple_dz/pkg/tadaptor"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router        fiber.Router
	cusstomLogger *zerolog.Logger
	repository    *vacancy.VacancyRepository
	store         *session.Store
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLoger *zerolog.Logger, repository *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:        router,
		cusstomLogger: customLoger,
		repository:    repository,
		store:         store,
	}

	h.router.Get("/", h.home)
	h.router.Get("/login", h.login)
	h.router.Get("/404", h.error)
	h.router.Post("/api/login", h.apiLogin)
	h.router.Get("/api/logout", h.apiLogout)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)

	count := int(math.Ceil(float64(h.repository.CountAll() / PAGE_ITEMS)))
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.cusstomLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, count, page)
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) apiLogin(c *fiber.Ctx) error {
	login := vacancy.LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if login.Email == "a@a.ru" && login.Password == "1" {
		session, err := h.store.Get(c)
		if err != nil {
			h.cusstomLogger.Error().Msg(err.Error())
			return c.SendStatus(500)
		}
		session.Set("email", login.Email)
		if err := session.Save(); err != nil {
			panic(err)
		}
		c.Response().Header.Add("HX-Redirect", "/")
		return c.Redirect("/", http.StatusOK)
	}
	component := components.Notification("Не верный email или пароль", components.NotificationFail)
	return tadaptor.Render(c, component, http.StatusBadRequest)
}

func (h *HomeHandler) apiLogout(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	session.Delete("email")
	if err := session.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("HX-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
