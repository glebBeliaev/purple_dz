package news

type News struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Content     string `db:"content"`
	Preview     string `db:"preview"`
	Category    string `db:"category"`
	Cover       string `db:"cover"`
	AuthorEmail string `db:"author_email"`
	CreatedAt   string `db:"created_at"`
}

type NewsForm struct {
	Title    string
	Content  string
	Preview  string
	Category string
	Cover    string
	Author   string
}
