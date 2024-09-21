package middlewares

import (
    "github.com/labstack/echo/v4"
)

func Cookie(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        accessToken, _ := c.Cookie("access_token")

        if accessToken != nil {
            c.Request().Header.Set("Authorization", "Bearer "+accessToken.Value)
        }

        return next(c)
    }
}
