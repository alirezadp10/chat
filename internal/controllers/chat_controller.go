package controllers

import (
    "encoding/json"
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/alirezadp10/chat/internal/mqtt"
    "github.com/alirezadp10/chat/pkg/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

type ChatParticipant struct {
    Name     string
    Username string
    Avatar   string
    ChatName string
}

type ChatMessage struct {
    models.Message
    models.User
}

type Message struct {
    Message  string `json:"message"`
    ClientID uint   `json:"clientId"`
}

func Chats(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.String(http.StatusNotFound, err.Error())
    }

    query := `
        SELECT u.name as name, u.username as username, u.avatar_url as avatar, c.name as chat_name
        FROM chat_participants
        JOIN chat.users u ON chat_participants.user_id = u.id
        JOIN chat.chats c ON c.id = chat_participants.chat_id
        WHERE chat_participants.chat_id IN (
            SELECT chat_id
            FROM chat_participants
            WHERE user_id = ?
        )
        AND chat_participants.user_id != ?
    `

    var chatParticipants []ChatParticipant

    db.Connection().Raw(query, user.ID, user.ID).Scan(&chatParticipants)

    var response []interface{}

    for _, participant := range chatParticipants {
        response = append(response, map[string]interface{}{
            "name":      participant.Name,
            "username":  participant.Username,
            "chat_name": participant.ChatName,
            "status":    "Online",
            "avatar":    "https://via.placeholder.com/50",
        })
    }

    return c.JSON(http.StatusOK, response)
}

func ShowChat(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.String(http.StatusNotFound, err.Error())
    }

    var messages []ChatMessage

    query := `
        SELECT * FROM messages join chat.chats c on c.id = messages.chat_id where c.name = ?;
    `

    db.Connection().Raw(query, c.Param("chatName")).Scan(&messages)

    var response []interface{}

    for _, message := range messages {
        isSelf := false
        if message.SenderID == user.ID {
            isSelf = true
        }
        response = append(response, map[string]interface{}{
            "message": message.Content,
            "is_self": isSelf,
        })
    }

    return c.JSON(http.StatusOK, response)
}

func SendMessage(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.String(http.StatusNotFound, err.Error())
    }

    messageData, _ := json.Marshal(Message{
        Message:  c.FormValue("message"),
        ClientID: user.ID,
    })

    mqtt.Client.Publish(c.Param("chatName"), 0, false, messageData)

    var chat models.Chat

    query := `SELECT id FROM chats where chats.name = ?;`

    db.Connection().Raw(query, c.Param("chatName")).Scan(&chat)

    newMessage := models.Message{
        Content:  c.FormValue("message"),
        ChatID:   chat.ID,
        SenderID: user.ID,
    }

    result := db.Connection().Create(&newMessage)

    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": result.Error.Error(),
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

func Search(c echo.Context) error {
    var users []models.User

    // Get the query parameter from the request
    searchQuery := c.QueryParam("query")

    // Build the SQL pattern string
    pattern := "%" + searchQuery + "%"

    // Define the query with a placeholder
    query := `
        SELECT * FROM users WHERE username LIKE ?;
    `

    // Execute the query with the pattern and handle potential errors
    if err := db.Connection().Raw(query, pattern).Scan(&users).Error; err != nil {
        // Return a 500 Internal Server Error if something goes wrong
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // Construct the response
    response := []interface{}{}
    for _, user := range users {
        response = append(response, map[string]interface{}{
            "name":     user.Name,
            "username": user.Username,
            "status":   "Online",
            "avatar":   "https://via.placeholder.com/50",
        })
    }

    // Return the JSON response
    return c.JSON(http.StatusOK, response)
}
