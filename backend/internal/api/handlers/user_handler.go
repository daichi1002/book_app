package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/daichi1002/book_app/backend/internal/services"
)

type UserHandler struct {
	Service services.UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
