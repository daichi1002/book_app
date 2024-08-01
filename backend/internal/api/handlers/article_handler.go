package handlers

import (
	"net/http"
	"strconv"

	"github.com/daichi1002/book_app/backend/internal/services"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	Service services.ArticleService
}

func (h *ArticleHandler) GetArticles(c echo.Context) error {
	articles, err := h.Service.GetArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticle(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	article, err := h.Service.GetArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, article)
}
