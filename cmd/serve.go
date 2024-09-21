package cmd

import (
    "github.com/alirezadp10/chat/internal/configs"
    "github.com/alirezadp10/chat/internal/controllers"
    "github.com/alirezadp10/chat/internal/middlewares"
    "github.com/alirezadp10/chat/internal/mqtt"
    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "run http server",
    Run:   serve,
}

func init() {
    rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
    mqtt.StartMQTT()

    e := echo.New()

    // TODO use pipeline
    e.Use(middleware.CORSWithConfig(configs.Cors()))

    // Public routes
    e.POST("/login", controllers.Login)
    e.POST("/register", controllers.Register)

    // Authenticated routes
    authGroup := e.Group("/api", middlewares.Cookie, echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(configs.JWT()["secret"].(string)),
    }))
    authGroup.GET("/users/search", controllers.Search, middlewares.Auth())
    authGroup.GET("/chats", controllers.Chats, middlewares.Auth())
    authGroup.GET("/chats/:username", controllers.ShowChat, middlewares.Auth())
    authGroup.POST("/chats/:chatName/messages", controllers.SendMessage, middlewares.Auth())

    e.Logger.Fatal(e.Start(configs.App()["url"].(string)))
}
