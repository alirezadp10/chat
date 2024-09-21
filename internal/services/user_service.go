package services

import (
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
)

type UserService struct{}

type User struct {
    User     string `json:"message"`
    ClientID uint   `json:"clientId"`
}

func NewUserService() *UserService {
    return &UserService{}
}

func (s *UserService) Search(searchQuery string) ([]interface{}, error) {
    response := []interface{}{}

    var users []models.User

    // Build the SQL pattern string
    pattern := "%" + searchQuery + "%"

    // Define the query with a placeholder
    query := `
        SELECT * FROM users WHERE username LIKE ?;
    `

    // Execute the query with the pattern and handle potential errors
    if err := db.Connection().Raw(query, pattern).Scan(&users).Error; err != nil {
        return response, err
    }

    // Construct the response
    for _, user := range users {
        response = append(response, map[string]interface{}{
            "id":       user.ID,
            "name":     user.Name,
            "username": user.Username,
            "status":   "Online",
            "avatar":   "https://via.placeholder.com/50",
        })
    }

    return response, nil
}
