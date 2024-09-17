package cmd

import (
    "github.com/alirezadp10/chat/internal/controllers"
    "github.com/alirezadp10/chat/internal/middlewares"
    "github.com/alirezadp10/chat/internal/mqtt"
    "github.com/labstack/echo/v4"
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

    // Public routes
    e.POST("/login", controllers.Login)
    e.POST("/register", controllers.Register)

    // Authenticated routes
    e.GET("/users/search", controllers.Search, middlewares.Auth())
    e.GET("/chats", controllers.Chats, middlewares.Auth())
    e.GET("/chats/:chatName", controllers.ShowChat, middlewares.Auth())
    e.POST("/chats/:chatId/messages", controllers.SendMessage, middlewares.Auth())

    e.Logger.Fatal(e.Start(":8080"))
}
