package main

// https://pkg.go.dev/github.com/tech-sumit/aws-iot-device-sdk-go#section-readme
// May need to work directly on paho to get all the capabilities

import (
	"fmt"
	"log"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
)

var TOPIC string = "ping"

// The ENDPOINT is kept in another file
var endpoint string = ENDPOINT

func Aws_iot_device_sdk_example() {

	connection, err := mqtt.NewConnection(mqtt.Config{
		KeyPath:  ".cert/private.pem",
		CertPath: ".cert/certificate.pem",
		CAPath:   ".cert/root.pem",
		ClientId: "ping_client",
		Endpoint: endpoint,
	})

	if err != nil {
		log.Panicln(err)
	}

	defer func() {
		connection.Disconnect()
	}()

	go func() {
		err := connection.SubscribeWithHandler(TOPIC, 0, func(client MQTT.Client, message MQTT.Message) {
			fmt.Println(string(message.Payload()))
		})

		if err != nil {
			log.Panicln(err)
		}
	}()

	go func() {
		for n := 0; n < 5; n++ {
			time.Sleep(1 * time.Second)

			err := connection.Publish(TOPIC, "pong - "+strconv.Itoa(n), 0)
			if err != nil {
				panic(err)
			}
		}
	}()

	fmt.Println("Press any key to stop")
	fmt.Println()

	fmt.Scanln()

}
