package models

const UserContextKey = "user"

type AuthUser struct {
	Email    string
	LoggedIn bool
}
