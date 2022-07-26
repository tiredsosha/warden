package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/tiredsosha/warden/mosquitto"
	"gopkg.in/yaml.v3"
)

type conf struct {
	Broker   string `yaml:"broker"`
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
		log.Println("Can not get hostname")
		hostname = "default"
	}
	log.Printf("Hostname: %s\n", hostname)
	return
}

func logInit() {
	file, err := os.OpenFile("warden.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.SetOutput(os.Stderr)
		log.Println("Can't open/make a warden.log. Logging in console")

	} else {
		log.SetOutput(file)
	}
}

func main() {
	logInit()

	log.Println("")
	log.Println("warden started")

	cfg := &conf{}
	if err := getConf("config.yaml", cfg); err != nil {
		log.Fatal("No config.yaml file in warden.exe file")
	}

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
