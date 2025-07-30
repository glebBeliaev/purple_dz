package models

import "github.com/glebbeliaev/purple_dz/internal/news"

type CategoryBlock struct {
	Title      string
	News       []news.News
	Page       int
	Count      int
	QueryParam string
}
