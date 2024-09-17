package mqtt

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
)

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

var Client MQTT.Client

func StartMQTT() MQTT.Client {
    // Create MQTT client options
    opts := MQTT.NewClientOptions()
    opts.AddBroker("tcp://test.mosquitto.org:1883")
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler

    // Create the client and connect
    Client = MQTT.NewClient(opts)
    if token := Client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    return Client
}
