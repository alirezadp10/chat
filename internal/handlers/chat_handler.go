package handlers

import (
    "github.com/alirezadp10/chat/internal/services"
    "github.com/alirezadp10/chat/pkg/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

type ChatHandler struct {
    Service *services.ChatService
}

func NewChatHandler(service *services.ChatService) *ChatHandler {
    return &ChatHandler{Service: service}
}

func (h *ChatHandler) Index(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.JSON(http.StatusUnauthorized, []interface{}{
            err.Error(),
        })
    }

    chatParticipants := h.Service.List(*user)

    var response []interface{}

    for _, participant := range chatParticipants {
        response = append(response, map[string]interface{}{
            "user_id":   participant.UserID,
            "name":      participant.Name,
            "username":  participant.Username,
            "chat_name": participant.ChatName,
            "status":    "Online",
            "avatar":    "https://via.placeholder.com/50",
        })
    }

    return c.JSON(http.StatusOK, response)
}

func (h *ChatHandler) Show(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.JSON(http.StatusUnauthorized, []interface{}{
            err.Error(),
        })
    }

    chatName, messages, err := h.Service.Show(c.Param("username"), *user)

    if err != nil {
        return c.JSON(http.StatusInternalServerError, []interface{}{
            err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "messages": messages,
        "chatName": chatName,
    })
}
