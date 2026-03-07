package items

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/MiKance/ToDoAPI/internal/service"
)

type ItemHandler struct {
	service.ToDoItem
}

func NewItemHandler(item service.ToDoItem) ItemHandler {
	return ItemHandler{item}
}

func (h *ItemHandler) GetItemByListId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(r.PathValue("user_id"))

		listID, err := strconv.Atoi(r.PathValue("list_id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		item, err := h.ToDoItem.GetItemsByListID(ctx, listID, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(item)
	}
}

func (h *ItemHandler) GetItemByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(r.PathValue("user_id"))

		itemID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		item, err := h.ToDoItem.GetItemsByID(ctx, itemID, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(item)
	}
}
