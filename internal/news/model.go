package news

import (
	"time"
)

type News struct {
	Id        int       `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Preview   string    `db:"preview"`
	Category  string    `db:"category"`
	Cover     string    `db:"cover"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"datacreate"`
}

type NewsForm struct {
	Title    string
	Content  string
	Preview  string
	Category string
	Cover    string
	Author   string
}
