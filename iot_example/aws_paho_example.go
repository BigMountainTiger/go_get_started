package main

// https://github.com/eclipse/paho.mqtt.golang
// https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang
// https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
// https://github.com/eclipse/paho.mqtt.golang/issues/72

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connected bool = false

// Make sense to use pointer here to make sure
// the same lock is aquired
var lock *sync.Mutex = &sync.Mutex{}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	topic := TOPIC
	token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})

	token.Wait()

	connected = true
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {

	connected = false
	fmt.Printf("Connect lost: %v\n", err)
}

func Aws_paho_example() {

	ca_path := ".cert/root.pem"
	key_path := ".cert/private.pem"
	cert_path := ".cert/certificate.pem"

	certpool := x509.NewCertPool()
	ca, err := os.ReadFile(ca_path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	certpool.AppendCertsFromPEM(ca)
	clientKeyPair, err := tls.LoadX509KeyPair(cert_path, key_path)
	if err != nil {
		panic(err)
	}

	tls_config := &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{clientKeyPair},
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcps://%s:%d/mqtt", ENDPOINT, 8883))
	opts.SetClientID("ping_client")
	opts.CleanSession = true
	opts.KeepAlive = 60
	opts.TLSConfig = tls_config

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {
		for {
			time.Sleep(5 * time.Second)

			if !connected {
				break
			}

			go func() {
				lock.Lock()
				defer lock.Unlock()

				token := client.Publish(TOPIC, 1, false, "Pong")
				token.Wait()
			}()
		}
	}()

	fmt.Print("Press 'Enter' to stop...")
	fmt.Println()
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// After calling disconnect, the client will not try to re-connect
	client.Disconnect(1000)
}
