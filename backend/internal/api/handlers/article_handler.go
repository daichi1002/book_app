package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/daichi1002/book_app/backend/internal/services"
)

type ArticleHandler struct {
	Service services.ArticleService
}

func (h *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.Service.GetArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articles)
}
