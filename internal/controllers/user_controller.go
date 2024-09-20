package controllers

import (
    "github.com/alirezadp10/chat/pkg/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

func Home(c echo.Context) error {
    user, err := utils.GetAuthUser(c)

    if err != nil {
        return c.JSON(http.StatusNotFound, []interface{}{
            err.Error(),
        })
    }

    return c.String(http.StatusOK, "Welcome "+user.Email+"!")
}
