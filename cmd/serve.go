package cmd

import (
    "github.com/alirezadp10/chat/internal/configs"
    "github.com/alirezadp10/chat/internal/handlers"
    "github.com/alirezadp10/chat/internal/middlewares"
    "github.com/alirezadp10/chat/internal/mqtt"
    "github.com/alirezadp10/chat/internal/services"
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

    // Initialize services
    chatService := services.NewChatService()
    messageService := services.NewMessageService()
    userService := services.NewUserService()

    // Initialize handlers
    chatHandler := handlers.NewChatHandler(chatService)
    messageHandler := handlers.NewMessageHandler(messageService)
    userHandler := handlers.NewUserHandler(userService)

    // Public routes
    e.POST("/login", handlers.Login)
    e.POST("/register", handlers.Register)

    // Authenticated routes
    authGroup := e.Group("/api", middlewares.Cookie, echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(configs.JWT()["secret"].(string)),
    }))
    authGroup.GET("/users/search", userHandler.Search, middlewares.Auth())
    authGroup.GET("/chats", chatHandler.Index, middlewares.Auth())
    authGroup.GET("/chats/:username", chatHandler.Show, middlewares.Auth())
    authGroup.POST("/chats/:chatName/messages", messageHandler.Send, middlewares.Auth())

    e.Logger.Fatal(e.Start(configs.App()["url"].(string)))
}
