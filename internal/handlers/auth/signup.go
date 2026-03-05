package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (h *AuthHandler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.User

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error(), "message": "invalid request body"})
			return
		}

		ctx := context.Background()
		ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		id, err := h.CreateUser(ctxTimeout, input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error(), "message": "server error"})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "message": "user created successfully"})
	}
}
