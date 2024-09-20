package models

import (
    "github.com/alirezadp10/chat/internal/db"
    "time"
)

type ChatParticipant struct {
    ID        uint `gorm:"primaryKey"`
    ChatID    uint `gorm:"not null"`
    UserID    uint `gorm:"not null"`
    UpdatedAt time.Time
    JoinedAt  time.Time
}

func AddParticipant(chatId uint, userId uint) ChatParticipant {
    newChatParticipant := ChatParticipant{
        ChatID: chatId,
        UserID: userId,
    }

    db.Connection().Create(&newChatParticipant)

    return newChatParticipant
}
