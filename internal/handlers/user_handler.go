package handlers

import (
    "github.com/alirezadp10/chat/internal/services"
    "github.com/labstack/echo/v4"
    "net/http"
)

type UserHandler struct {
    Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{Service: service}
}

func (h *UserHandler) Search(c echo.Context) error {
    //TODO use elastic or redis
    response, err := h.Service.Search(c.QueryParam("query"))

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // Return the JSON response
    return c.JSON(http.StatusOK, response)
}
