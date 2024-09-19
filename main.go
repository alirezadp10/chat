package main

import (
    "github.com/alirezadp10/chat/cmd"
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
)

func main() {
    _ = db.Connection().AutoMigrate(&models.User{}, &models.Chat{}, &models.Message{}, &models.ChatParticipant{})

    cmd.Execute()
}
