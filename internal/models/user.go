package models

import (
    "time"
)

type User struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null;size:100"`
    Password  string `gorm:"not null;size:100"`
    AvatarURL string
    Messages  []Message         `gorm:"foreignKey:SenderID"`
    Chats     []ChatParticipant `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
