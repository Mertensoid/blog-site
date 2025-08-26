package register

import "time"

type RegisterData struct {
	Name     string
	Email    string
	Password string
}

type LoginForm struct {
	Email    string
	Password string
}

type User struct {
	Id       int       `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"pass"`
	Name     string    `db:"name"`
	RegTime  time.Time `db:"regtime"`
}
