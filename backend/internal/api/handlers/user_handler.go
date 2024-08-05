package handlers

import (
	"net/http"

	"github.com/daichi1002/book_app/backend/internal/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service services.UserService
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.QueryParam("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
