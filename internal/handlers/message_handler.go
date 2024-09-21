package handlers

import (
    "github.com/alirezadp10/chat/internal/services"
    "github.com/alirezadp10/chat/pkg/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

type MessageHandler struct {
    Service *services.MessageService
}

func NewMessageHandler(service *services.MessageService) *MessageHandler {
    return &MessageHandler{Service: service}
}

func (h *MessageHandler) Send(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.JSON(http.StatusUnauthorized, []interface{}{
            err.Error(),
        })
    }

    err = h.Service.Send(*user, c.FormValue("message"), c.Param("chatName"))

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": err.Error(),
        })
    }

    //Disconnect the client
    //client.Disconnect(250)
    // Unsubscribe from the topic
    //if token := client.Unsubscribe(c.Param("chatId")); token.Wait() && token.Error() != nil {
    //   fmt.Println(token.Error())
    //   return c.String(http.StatusOK, c.FormValue("foo"))
    //}

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "success",
    })
}
