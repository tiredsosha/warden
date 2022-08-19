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
		logger.Warn.Printf("can't get hostname - %s\n", err)
		hostname = "default"
	}
	return
}

func main() {
	cfg, state, debug := config.ArgsInit()

	logger.LogInit(debug)
	logger.Debug.Println("")
	logger.Debug.Println("")
	logger.Info.Print("WARDENER STARTED")

	if !state {
		cfg = config.ConfInit()
	}

	hostname := getHostname()
	topicPrefix := "warden/" + hostname + "/"

	logger.Debug.Println("---------------------------")
	logger.Debug.Println("logging data:")
	logger.Debug.Printf("\tdebug    - %t\n", debug)
	logger.Debug.Printf("\tcli conf - %t\n", state)
	logger.Debug.Println("- - - - - - - - - - - - - -")
	logger.Debug.Println("сonnection data:")
	logger.Debug.Printf("\thostname - %s\n", hostname)
	logger.Debug.Printf("\tbroker   - %s\n", cfg.Broker)
	logger.Debug.Printf("\tusername - %s\n", cfg.Username)
	logger.Debug.Printf("\tpassword - %s\n", cfg.Password)
	logger.Debug.Println("---------------------------")

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
