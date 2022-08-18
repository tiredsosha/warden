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
	return
}

func main() {
	logger.Info.Println("")
	logger.Info.Println("")
	logger.Info.Print("WARDENER STARTED")

	cfg, state := config.CmdInit()

	if !state {
		cfg = config.ConfInit()
	}

	hostname := getHostname()
	topicPrefix := "warden/" + hostname + "/"

	logger.Info.Println("---------------------------")
	logger.Info.Println("сonnection data:")
	logger.Info.Printf("\targs     - %t\n", state)
	logger.Info.Printf("\thostname - %s\n", hostname)
	logger.Info.Printf("\tbroker   - %s\n", cfg.Broker)
	logger.Info.Printf("\tusername - %s\n", cfg.Username)
	logger.Info.Printf("\tpassword - %s\n", cfg.Password)
	logger.Info.Println("---------------------------")

	mqttData := mosquitto.MqttConf{
		ID:       hostname,
		Broker:   cfg.Broker,
		Username: cfg.Username,
		Password: cfg.Password,
		SubTopic: topicPrefix + "commands/",
		PubTopic: topicPrefix + "status/",
	}
	mosquitto.StartBroker(mqttData)
}
