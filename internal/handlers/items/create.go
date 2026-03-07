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

func (h *ItemHandler) CreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(r.PathValue("user_id"))
		listId, err := strconv.Atoi(chi.URLParam(r, "list_id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var input models.ToDoItem
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		input.ListID = listId

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		id, err := h.ToDoItem.CreateItem(ctx, input, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "message": "item created successfully"})
	}
}
