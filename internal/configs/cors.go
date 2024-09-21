package configs

import "github.com/labstack/echo/v4/middleware"

func Cors() middleware.CORSConfig {
    allowOrigins := []string{"*"}

    if App()["env"] != "production" {
        allowOrigins = []string{}
        allowOrigins = append(allowOrigins, "http://localhost:63342")
    }

    allowMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}

    return middleware.CORSConfig{
        AllowOrigins:     allowOrigins,
        AllowMethods:     allowMethods,
        AllowCredentials: true,
    }
}
