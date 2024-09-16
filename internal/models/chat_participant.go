package models

import (
    "time"
)

type ChatParticipant struct {
    ID        uint `gorm:"primaryKey"`
    ChatID    uint `gorm:"not null"`
    UserID    uint `gorm:"not null"`
    UpdatedAt time.Time
    JoinedAt  time.Time
}
