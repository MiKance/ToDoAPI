package auth

import (
	"encoding/json"
	"net/http"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error(), "message": "invalid request body"})
		}
	}
}
