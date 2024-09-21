package handlers

import (
    "errors"
    "github.com/alirezadp10/chat/internal/configs"
    "github.com/alirezadp10/chat/internal/db"
    "github.com/alirezadp10/chat/internal/form_requests"
    "github.com/alirezadp10/chat/internal/models"
    "github.com/alirezadp10/chat/pkg/utils"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "net/http"
    "time"
)

func Login(c echo.Context) error {
    userReq, err := form_requests.LoginFormRequest(c)

    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": err.Error(),
        })
    }

    var user models.User

    result := db.Connection().Where("username = ?", userReq.Username).Find(&user)

    if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": "Username or password is incorrect.",
        })
    }

    if !utils.Verify(userReq.Password, user.Password) {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": "Username or password is incorrect.",
        })
    }

    token, _ := utils.GenerateJWT(user.Username)

    setCookie(c, token)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  "success",
        "message": "User logged in successfully",
        "data": map[string]interface{}{
            "client_id":    user.ID,
            "access_token": token.AccessToken,
            "expire_at":    token.ExpireAt,
        },
    })
}

func setCookie(c echo.Context, token utils.Token) {
    cookie := new(http.Cookie)
    cookie.Name = "access_token"
    cookie.Value = token.AccessToken
    cookie.HttpOnly = true
    cookie.Secure = configs.Cookie()["secure"].(bool)
    cookie.Path = "/"
    expireAt, _ := time.Parse("2006-01-02 15:04:05", token.ExpireAt)
    cookie.Expires = expireAt
    c.SetCookie(cookie)
}

func Register(c echo.Context) error {
    //TODO sanitize input

    newUser, err := form_requests.RegisterFormRequest(c)

    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": err.Error(),
        })
    }

    result := db.Connection().Create(&newUser)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": result.Error.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  "success",
        "message": "User registered successfully",
        "data": map[string]interface{}{
            "id":         newUser.ID,
            "name":       newUser.Name,
            "username":   newUser.Username,
            "email":      newUser.Email,
            "created_at": newUser.CreatedAt,
            "updated_at": newUser.UpdatedAt,
        },
    })
}
