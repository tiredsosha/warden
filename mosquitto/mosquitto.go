package mosquitto

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tiredsosha/warden/control/power"
	"github.com/tiredsosha/warden/control/sound"
)

const (
	port          = 1883
	KeyLifeTime   = 2  // minute
	StateInterval = 25 // second
)

type MqttConf struct {
	Id       string
	Broker   string
	Username string
	Password string
	SubTopic string
	PubTopic string
}

func (data *MqttConf) MessageHandler(client mqtt.Client, msgHand mqtt.Message) {
	topic := msgHand.Topic()
	msg := strings.TrimSpace(string(msgHand.Payload()))
	executor(topic, msg, data.SubTopic)

}

func StartBroker(data MqttConf) {
	var wg sync.WaitGroup
	messagePubHandler := data.MessageHandler

	mqttHandler := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", data.Broker, port)).
		SetClientID(data.Id).
		SetUsername(data.Username).
		SetPassword(data.Password).
		SetAutoReconnect(false).
		SetDefaultPublishHandler(messagePubHandler).
		SetConnectionLostHandler(connectLostHandler).
		SetOnConnectHandler(connectHandler).
		SetKeepAlive(KeyLifeTime)

	conn := mqtt.NewClient(mqttHandler)
	if token := conn.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Can't connect to mqtt broker")
	}

	wg.Add(2)
	go subscribe(conn, data)
	go publish(conn, data)
	wg.Wait()
}

func publish(client mqtt.Client, data MqttConf) {
	i := 0
	for i < 1 {
		volume, err := sound.GetVolume()
		if err != nil {
			log.Println("Skiping 1 cycle of publishing")
		} else {
			strVolume := strconv.Itoa(int(volume))
			token := client.Publish(data.PubTopic+"volume", 0, false, strVolume)
			token.Wait()
		}
		time.Sleep(StateInterval * time.Second)
	}
}

func subscribe(client mqtt.Client, data MqttConf) {
	topic := data.SubTopic + "#"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic: %s\n", topic)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected to mqtt broker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Fatal("Connection to mqtt broker is lost")
}

func executor(topic, msg, subPrefix string) {
	log.Printf("%s recieved in %q\n", msg, topic)
	switch topic {
	case subPrefix + "volume":
		intMsg, err := strconv.Atoi(msg)
		if err == nil {
			sound.SetVolume(intMsg)
		} else {
			log.Println("Message must be in range of 0-100, skiping command")
		}
	case subPrefix + "mute":
		boolMsg, err := strconv.ParseBool(msg)
		if err == nil {
			sound.Mute(boolMsg)
		} else {
			log.Println("Message must be true or false, skiping command")
		}
	case subPrefix + "shutdown":
		power.Shutdown()
	case subPrefix + "reboot":
		power.Reboot()
	}
}
