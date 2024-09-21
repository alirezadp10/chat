package services

import (
    "encoding/json"
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/alirezadp10/chat/internal/mqtt"
)

type MessageService struct{}

type Message struct {
    Message  string `json:"message"`
    ClientID uint   `json:"clientId"`
}

func NewMessageService() *MessageService {
    return &MessageService{}
}

func (s *MessageService) Send(user models.User, message, chatName string) error {
    //TODO sanitize input
    messageData, _ := json.Marshal(Message{
        Message:  message,
        ClientID: user.ID,
    })

    mqtt.Client.Publish(chatName, 0, false, messageData)

    var chat models.Chat

    query := `SELECT id FROM chats where chats.name = ?;`

    db.Connection().Raw(query, chatName).Scan(&chat)

    newMessage := models.Message{
        Content:  message,
        ChatID:   chat.ID,
        SenderID: user.ID,
    }

    result := db.Connection().Create(&newMessage)

    if result.Error != nil {
        return result.Error
    }

    return nil
}
