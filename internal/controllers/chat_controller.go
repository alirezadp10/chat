package controllers

import (
    "fmt"
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/alirezadp10/chat/pkg/utils"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "github.com/labstack/echo/v4"
    "net/http"
)

// Message handler for received messages
var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// Message handler for successful connection
var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
    fmt.Println("Connected to MQTT broker")
}

// Handler for connection lost
var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
    fmt.Printf("Connection lost: %v\n", err)
}

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

    query := `
        SELECT * FROM messages join chat.chats c on c.id = messages.chat_id where c.name = ?;
    `

    var messages []ChatMessage

    fmt.Println(c.Param("chatName"))

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
    // Create MQTT client options
    opts := MQTT.NewClientOptions()
    opts.AddBroker("tcp://127.0.0.1:1883")
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler

    // Create the client and connect
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    text := c.FormValue("message")
    token := client.Publish(c.Param("chatId"), 0, false, text)
    token.Wait()

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "success",
    })
    // Unsubscribe from the topic
    //if token := client.Unsubscribe(c.Param("chatId")); token.Wait() && token.Error() != nil {
    //    fmt.Println(token.Error())
    //    return c.String(http.StatusOK, c.FormValue("foo"))
    //}

    // Disconnect the client
    //client.Disconnect(250)
}
