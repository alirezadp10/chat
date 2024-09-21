package middlewares

import (
    "errors"
    "github.com/labstack/echo/v4"
    "net/http"
)

func Cookie(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie("access_token")

        if err != nil {
            if errors.Is(err, http.ErrNoCookie) {
                return echo.NewHTTPError(http.StatusUnauthorized, "Cookie not found")
            }
            return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving cookie")
        }

        c.Request().Header.Set("Authorization", "Bearer "+cookie.Value)

        return next(c)
    }
}
