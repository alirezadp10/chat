package cmd

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "github.com/labstack/echo/v4"
    "github.com/spf13/cobra"
    "net/http"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "run http server",
    Run:   serve,
}

func init() {
    rootCmd.AddCommand(serveCmd)
}

// Message handler for received messages
var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// Message handler for successful connection
var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
    fmt.Println("Connected to MQTT broker")
}

// Handler for connection lost
var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
    fmt.Printf("Connection lost: %v\n", err)
}

func serve(cmd *cobra.Command, args []string) {
    e := echo.New()

    e.GET("/chats", func(c echo.Context) error {
        return c.JSON(http.StatusOK, []interface{}{
            map[string]interface{}{
                "name":     "User 1",
                "username": "User 1",
                "chatId":   "1234",
                "status":   "Online",
                "avatar":   "https://via.placeholder.com/50",
            },
            map[string]interface{}{
                "name":     "User 2",
                "username": "User 2",
                "chatId":   "4312",
                "status":   "Online",
                "avatar":   "https://via.placeholder.com/50",
            },
            map[string]interface{}{
                "name":     "User 3",
                "username": "User 3",
                "chatId":   "5432",
                "status":   "Last Seen Recently",
                "avatar":   "https://via.placeholder.com/50",
            },
            map[string]interface{}{
                "name":     "User 4",
                "username": "User 4",
                "chatId":   "5132",
                "status":   "Last Seen Recently",
                "avatar":   "https://via.placeholder.com/50",
            },
            map[string]interface{}{
                "name":     "User 5",
                "username": "User 5",
                "chatId":   "1432",
                "status":   "Last Seen Recently",
                "avatar":   "https://via.placeholder.com/50",
            },
        })
    })

    e.GET("/chats/:chatId", func(c echo.Context) error {
        return c.JSON(http.StatusOK, []interface{}{
            map[string]interface{}{
                "message": "hi",
                "isSelf":  false,
            },
            map[string]interface{}{
                "message": "hello",
                "isSelf":  true,
            },
        })
    })

    e.POST("/chats/:chatId/messages", func(c echo.Context) error {
        // Create MQTT client options
        opts := MQTT.NewClientOptions()
        opts.AddBroker("tcp://127.0.0.1:1883")
        opts.SetDefaultPublishHandler(messagePubHandler)
        opts.OnConnect = connectHandler
        opts.OnConnectionLost = connectLostHandler

        // Create the client and connect
        client := MQTT.NewClient(opts)
        if token := client.Connect(); token.Wait() && token.Error() != nil {
            panic(token.Error())
        }

        text := c.FormValue("message")
        token := client.Publish(c.Param("chatId"), 0, false, text)
        token.Wait()

        return c.JSON(http.StatusOK, map[string]interface{}{
            "status": "success",
        })
        // Unsubscribe from the topic
        //if token := client.Unsubscribe(c.Param("chatId")); token.Wait() && token.Error() != nil {
        //    fmt.Println(token.Error())
        //    return c.String(http.StatusOK, c.FormValue("foo"))
        //}

        // Disconnect the client
        //client.Disconnect(250)
    })

    e.Logger.Fatal(e.Start(":8080"))
}
