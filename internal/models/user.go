package models

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
