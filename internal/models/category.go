package models

import (
	"math"

	"github.com/glebbeliaev/purple_dz/internal/news"
	"github.com/gofiber/fiber/v2"
)

type CategoryBlock struct {
	Title      string
	News       []news.News
	Page       int
	Count      int
	QueryParam string
}

func BuildCategoryBlock(c *fiber.Ctx, repo *news.NewsRepository, category, title, queryParam string, limit int) (CategoryBlock, error) {
	page := c.QueryInt(queryParam, 1)

	var total int
	var newsItems []news.News
	var err error

	if category == "" {
		total = repo.CountAll()
		newsItems, err = repo.GetAll(limit, (page-1)*limit)
	} else {
		total = repo.CountByCategory(category)
		newsItems, err = repo.GetByCategory(category, limit, (page-1)*limit)
	}

	if err != nil {
		return CategoryBlock{}, err
	}

	count := int(math.Ceil(float64(total) / float64(limit)))

	return CategoryBlock{
		Title:      title,
		News:       newsItems,
		Page:       page,
		Count:      count,
		QueryParam: queryParam,
	}, nil
}

type CategoryButton struct {
	Tag  string
	Link string
	Img  string
}

var AllCategories = []CategoryButton{
	{Tag: "#Еда", Link: "/food", Img: "/public/images/categories/categoryFood.png"},
	{Tag: "#Животные", Link: "/animal", Img: "/public/images/categories/categoryAnimal.png"},
	{Tag: "#Машины", Link: "/cars", Img: "/public/images/categories/categoryCars.png"},
	{Tag: "#Спорт", Link: "/sport", Img: "/public/images/categories/categorySport.png"},
	{Tag: "#Музыка", Link: "/music", Img: "/public/images/categories/categoryMusic.png"},
	{Tag: "#Технологии", Link: "/tech", Img: "/public/images/categories/categoryTechnology.png"},
	{Tag: "#Прочее", Link: "/other", Img: "/public/images/categories/categoryMore.png"},
}
