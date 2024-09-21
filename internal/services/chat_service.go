package services

import (
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/alirezadp10/chat/pkg/utils"
    "gorm.io/gorm"
)

type ChatService struct{}

type ChatMessages struct {
    models.Message
    models.User
}

func NewChatService() *ChatService {
    return &ChatService{}
}

func (s *ChatService) List(user models.User) []struct {
    UserID   uint
    Name     string
    Username string
    Avatar   string
    ChatName string
} {
    query := `
        SELECT u.id as user_id, u.name as name, u.username as username, u.avatar_url as avatar, c.name as chat_name
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

    var chatParticipants []struct {
        UserID   uint
        Name     string
        Username string
        Avatar   string
        ChatName string
    }

    db.Connection().Raw(query, user.ID, user.ID).Scan(&chatParticipants)

    return chatParticipants
}

func (s *ChatService) Show(username string, user models.User) (string, []ChatMessages, error) {
    var audience models.User

    db.Connection().Where("username = ?", username).Find(&audience)
    //-------------------------------------------------------------------------------------
    var chatID struct{ ID uint }

    query := `
        SELECT chat_id as id FROM chat_participants WHERE user_id in (?,?) GROUP BY chat_id HAVING COUNT(DISTINCT user_id) = 2
    `

    db.Connection().Raw(query, audience.ID, user.ID).Scan(&chatID)

    var chat models.Chat

    query = `
        SELECT * FROM chats WHERE id = ?
    `

    result := db.Connection().Raw(query, chatID.ID).Scan(&chat)

    chatName := chat.Name

    var chatMessages []ChatMessages

    if result.RowsAffected <= 0 {
        err := db.Connection().Transaction(func(tx *gorm.DB) error {
            // Generate chat name
            chatName, err := utils.RandomString(10)
            if err != nil {
                return err // Return error to rollback
            }

            // Create new chat
            newChat := models.Chat{Name: chatName}
            if err = tx.Create(&newChat).Error; err != nil {
                return err // Return error to rollback
            }

            // Add participants
            newChatParticipant := models.ChatParticipant{ChatID: newChat.ID, UserID: audience.ID}
            if err = tx.Create(&newChatParticipant).Error; err != nil {
                return err // Return error to rollback
            }

            newChatParticipant = models.ChatParticipant{ChatID: newChat.ID, UserID: user.ID}
            if err = tx.Create(&newChatParticipant).Error; err != nil {
                return err // Return error to rollback
            }

            return nil // Return nil to commit the transaction
        })

        if err != nil {
            // Handle the error (transaction is rolled back)
            //c.Param("username")
            return "", nil, err
        }
    } else {
        query = `
            SELECT * FROM messages join chat.chats c on c.id = messages.chat_id where c.id = ?;
        `
        db.Connection().Raw(query, chat.ID).Scan(&chatMessages)
    }

    msgs := []interface{}{}

    for _, message := range chatMessages {
        isSelf := false
        if message.SenderID == user.ID {
            isSelf = true
        }
        msgs = append(msgs, map[string]interface{}{
            "message": message.Content,
            "is_self": isSelf,
        })
    }

    return chatName, chatMessages, nil
}