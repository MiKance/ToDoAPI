package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (h *AuthHandler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.UserSignIn

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error(), "message": "invalid request body"})
			return
		}

		ctx := context.Background()
		ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		token, err := h.GenerateToken(ctxTimeout, input.Username, input.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error(), "message": "server error"})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"token": token, "message": "jwt created successfully"})
	}
}
