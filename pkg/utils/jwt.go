package utils

import (
    "errors"
    "fmt"
    "github.com/alirezadp10/chat/internal/configs"
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "strconv"
    "time"
)

type Token struct {
    AccessToken string
    ExpireAt    string
}

func GenerateJWT(userID string) (Token, error) {
    jwtSecret := []byte(configs.JWT()["secret"])

    tokenLifeTime, _ := strconv.Atoi(configs.JWT()["tokenLifeTime"])

    claims := jwt.MapClaims{
        "name": userID,
        "exp":  time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtSecret)

    if err != nil {
        return Token{}, err
    }

    return Token{
        AccessToken: tokenString,
        ExpireAt:    time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Format("2006-01-02 15:04:05"),
    }, nil
}

func GetAuthUser(c echo.Context) (*models.User, error) {
    var user models.User
    claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    result := db.Connection().Where("username = ?", name).Find(&user)

    if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, fmt.Errorf("failed to find authenticated user")
    }

    return &user, nil
}
