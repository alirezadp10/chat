package models

import "time"

type Message struct {
    ID        uint   `gorm:"primaryKey"`
    ChatID    uint   `gorm:"not null"`
    SenderID  uint   `gorm:"not null"`
    Content   string `gorm:"type:text;not null"` // message_content
    IsRead    bool   `gorm:"default:false"`
    CreatedAt time.Time
}
