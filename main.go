package main

import (
	"os"

	"github.com/tiredsosha/warden/mosquitto"
	config "github.com/tiredsosha/warden/tools/configurator"
	"github.com/tiredsosha/warden/tools/logger"
)

func getHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Warn.Println("can't get hostname")
		logger.Warn.Println(err)
		hostname = "default"
	}
	logger.Info.Printf("hostname is %s\n", hostname)
	return
}

func main() {
	logger.LogInit()

	logger.Info.Println("")
	logger.Info.Println("WARDEN STARTED")

	cfg := config.ConfInit()

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
