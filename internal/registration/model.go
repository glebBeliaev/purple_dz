package registration

import "time"

type RegistrationForm struct {
	Name     string
	Email    string
	Password string
}

type User struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Pass      string    `db:"pass"`
	UserName  string    `db:"username"`
	CreatedAt time.Time `db:"createdat"`
}
