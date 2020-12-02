package main

import (
	applications "bridge/src/applications"
	frameworks "bridge/src/frameworks"
)

func main() {
	messaging := frameworks.NewMessaging()
	err := messaging.Connect()
	if err != nil {
		panic("RabbitMQ Connection Err")
	}

	mqttClient := frameworks.NewMqttClient()
	err = mqttClient.ConnectToMQTT()
	if err != nil {
		panic("MQTT Connection Err")
	}

	bridge := applications.NewSendToMessagingService(mqttClient, messaging)
	bridge.Usecase()
}
