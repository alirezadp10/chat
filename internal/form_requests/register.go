package form_requests

import (
    "errors"
    "fmt"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/alirezadp10/chat/pkg/utils"
    "github.com/labstack/echo/v4"
)

func RegisterFormRequest(c echo.Context) (models.User, error) {
    var userReq models.User

    // Decode JSON body and handle errors
    if err := c.Bind(&userReq); err != nil {
        return models.User{}, fmt.Errorf("failed to decode request body: %w", err)
    }

    // Hash the password and handle errors
    hash, err := utils.Hash(userReq.Password)
    if err != nil {
        return models.User{}, fmt.Errorf("failed to hash password: %w", err)
    }

    // Validate the user request
    if err := validateRegisterForm(userReq); err != nil {
        return models.User{}, fmt.Errorf("validation failed: %w", err)
    }

    // Create new user with hashed password
    newUser := models.User{
        Name:     userReq.Name,
        Username: userReq.Username,
        Email:    userReq.Email,
        Password: hash,
    }

    return newUser, nil
}

func validateRegisterForm(u models.User) error {
    if u.Name == "" || u.Username == "" || u.Password == "" || u.Email == "" {
        return errors.New("missing required fields")
    }
    return nil
}
