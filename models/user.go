package models

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
