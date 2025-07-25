package news

type News struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Content     string `db:"content"`
	AuthorEmail string `db:"author_email"`
	CreatedAt   string `db:"created_at"`
}
