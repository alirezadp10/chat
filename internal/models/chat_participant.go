package models

import (
    "github.com/alirezadp10/chat/internal/db"
    "gorm.io/gorm"
    "time"
)

type ChatParticipant struct {
    ID        uint `gorm:"primaryKey"`
    ChatID    uint `gorm:"not null"`
    UserID    uint `gorm:"not null"`
    UpdatedAt time.Time
    JoinedAt  time.Time
}

func AddParticipant(tx *gorm.DB, chatId, userId uint) ChatParticipant {
    newChatParticipant := ChatParticipant{
        ChatID: chatId,
        UserID: userId,
    }

    db.Connection().Create(&newChatParticipant)

    return newChatParticipant
}
