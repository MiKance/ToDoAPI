package items

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/go-chi/chi/v5"
)

func (h *ItemHandler) UpdateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(r.PathValue("user_id"))

		itemID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "id not found"})
			return
		}

		var input models.ToDoItem
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		input.ID = itemID

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = h.ToDoItem.UpdateItem(ctx, input, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "item updated successfully"})

	}
}
