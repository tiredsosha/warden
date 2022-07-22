package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tiredsosha/wardener/mosquitto"
	"gopkg.in/yaml.v3"
)

type conf struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getConf(file string, cnf interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, cnf)
	}
	return err
}

func getHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n", hostname)
	return
}

func main() {
	cfg := &conf{}
	if err := getConf("config.yaml", cfg); err != nil {
		log.Panicln(err)
	}

	hostname := getHostname()
	mqttData := mosquitto.MqttConf{
		Id:       hostname,
		Broker:   cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
		SubTopic: "wardener/#",
		PubTopic: "wardener/sound/status",
	}
	mosquitto.StartBroker(mqttData)
}
