package middlewares

import (
    "github.com/alirezadp10/chat/internal/configs"
    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
    return echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(configs.JWT()["secret"]),
    })
}
