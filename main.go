package main

import (
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
		log.Println("Can not get hostname")
		hostname = "default"
	}
	log.Printf("Hostname: %s\n", hostname)
	return
}

func main() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("wardener.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Wardener started")

	cfg := &conf{}
	if err := getConf("config.yaml", cfg); err != nil {
		log.Fatal("No config.yaml file in wardener.exe file")
	}

	hostname := getHostname()
	mqttData := mosquitto.MqttConf{
		Id:       hostname,
		Broker:   cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
	}
	mosquitto.StartBroker(mqttData)
}
