package main

import (
	"os"

	"github.com/tiredsosha/warden/mosquitto"
	config "github.com/tiredsosha/warden/tools/configurator"
	"github.com/tiredsosha/warden/tools/logger"
	"github.com/tiredsosha/warden/tray"
)

func getHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Warn.Printf("can't get hostname - %s\n", err)
		hostname = "default"
	}
	return
}

func main() {
	cfg, state, debug := config.ArgsInit()

	logger.LogInit(debug)

	if !state {
		cfg = config.ConfInit()
	}

	hostname := getHostname()
	topicPrefix := "warden/" + hostname + "/"

	logger.DebugLog(debug, state, hostname, cfg.Broker, cfg.Username, cfg.Password)

	mqttData := mosquitto.MqttConf{
		ID:       hostname,
		Broker:   cfg.Broker,
		Username: cfg.Username,
		Password: cfg.Password,
		SubTopic: topicPrefix + "commands/",
		PubTopic: topicPrefix + "status/",
	}
	go tray.TrayStart()
	mosquitto.StartBroker(mqttData)
}
