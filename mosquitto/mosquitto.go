package mosquitto

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tiredsosha/wardener/control/power"
	"github.com/tiredsosha/wardener/control/sound"
)

type MqttConf struct {
	Id       string
	Broker   string
	Username string
	Password string
}

const (
	port          = 1883
	KeyLifeTime   = 2  // minute
	StateInterval = 10 // second
	SubTopic      = "wardener/command/"
	PubTopic      = "wardener/status/"
)

var wg sync.WaitGroup

func StartBroker(data MqttConf) {
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
		panic(token.Error())
	}

	wg.Add(2)
	go subscribe(conn, data)
	go publish(conn, data)
	wg.Wait()
	// conn.Disconnect(250)
}

func publish(client mqtt.Client, data MqttConf) {
	i := 0
	for i < 1 {
		volume, err := sound.GetVolume()
		if err == nil {
			strVolume := strconv.Itoa(int(volume))
			// fmt.Println(strVolume, err)

			token := client.Publish(PubTopic+"volume", 0, false, strVolume)
			token.Wait()
		}
		time.Sleep(time.Second)
	}
}

func subscribe(client mqtt.Client, data MqttConf) {
	topic := SubTopic + "#"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msgHand mqtt.Message) {
	topic := msgHand.Topic()
	msg := strings.TrimSpace(string(msgHand.Payload()))
	executor(topic, msg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func executor(topic, msg string) {
	switch topic {
	case SubTopic + "volume":
		intMsg, err := strconv.Atoi(msg)
		if err == nil {
			sound.SetVolume(intMsg)
		}
	case SubTopic + "mute":
		boolMsg, err := strconv.ParseBool(msg)
		if err == nil {
			sound.Mute(boolMsg)
		}
	case SubTopic + "shutdown":
		power.Shutdown()
	case SubTopic + "reboot":
		power.Reboot()
	default:
		fmt.Printf("%s recieved in %d\n", msg, topic)
	}
}
