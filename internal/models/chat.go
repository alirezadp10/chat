package models

import (
    "time"
)

type Chat struct {
    ID           uint              `gorm:"primaryKey"`
    Name         string            // Optional, e.g. for group chats
    Participants []ChatParticipant `gorm:"foreignKey:ChatID"`
    Messages     []Message         `gorm:"foreignKey:ChatID"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
