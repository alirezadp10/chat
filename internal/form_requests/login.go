package form_requests

import (
    "errors"
    "fmt"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/labstack/echo/v4"
)

func LoginFormRequest(c echo.Context) (models.User, error) {
    var userReq models.User

    if err := c.Bind(&userReq); err != nil {
        return models.User{}, fmt.Errorf("failed to decode request body: %w", err)
    }

    // Validate the user request
    if err := validateLoginForm(userReq); err != nil {
        return models.User{}, fmt.Errorf("validation failed: %w", err)
    }

    return userReq, nil
}

func validateLoginForm(u models.User) error {
    if u.Username == "" || u.Password == "" {
        return errors.New("missing required fields")
    }
    return nil
}
