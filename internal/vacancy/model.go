package vacancy

import "time"

type VacancyCreateFrom struct {
	Email       string
	Role        string
	Company     string
	CompanyType string
	Salary      string
	Location    string
}

type Vacancy struct {
	Id          int       `db:"id"`
	Email       string    `db:"email"`
	Role        string    `db:"role"`
	Company     string    `db:"company"`
	CompanyType string    `db:"type"`
	Salary      string    `db:"salary"`
	Location    string    `db:"location"`
	CreatedAt   time.Time `db:"createdat"`
}

type LoginForm struct {
	Email    string
	Password string
}
