package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) userIdentify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Authorization header is empty"})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Authorization header is incorrect"})
			return
		}

		userID, err := h.service.ParseToken(headerParts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
			return
		}
		r.SetPathValue("user_id", strconv.Itoa(userID))
		next.ServeHTTP(w, r)
	})
}
