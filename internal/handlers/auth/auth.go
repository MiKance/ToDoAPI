package auth

import "github.com/MiKance/ToDoAPI/internal/service"

type AuthHandler struct {
	service.Authorization
}

func NewAuthHandler(auth service.Authorization) *AuthHandler {
	return &AuthHandler{auth}
}
