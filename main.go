package main

import (
	"log"
	"os"

	"github.com/tiredsosha/warden/mosquitto"
)

func getHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("can't get hostname")
		hostname = "default"
	}
	log.Printf("hostname: %s\n", hostname)
	return
}

func main() {
	LogInit()
	cfg := ConfInit()

	log.Println("")
	log.Println("WARDEN STARTED")

	hostname := getHostname()
	topicPrefix := "warden/" + hostname + "/"

	mqttData := mosquitto.MqttConf{
		Id:       hostname,
		Broker:   cfg.Broker,
		Username: cfg.Username,
		Password: cfg.Password,
		SubTopic: topicPrefix + "command/",
		PubTopic: topicPrefix + "status/",
	}
	mosquitto.StartBroker(mqttData)
}

func ConfInit() {
	panic("unimplemented")
}
