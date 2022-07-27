package mosquitto

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tiredsosha/warden/control/power"
	"github.com/tiredsosha/warden/control/sound"
	"github.com/tiredsosha/warden/tools/logger"
)

const (
	port          = 1883
	KeyLifeTime   = 2  // minute
	StateInterval = 25 // second
)

type pubFunc func() (string, error)

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
		SetKeepAlive(KeyLifeTime).
		SetWill(data.PubTopic+"online", "false", 2, true)

	conn := mqtt.NewClient(mqttHandler)
	if token := conn.Connect(); token.Wait() && token.Error() != nil {
		logger.Error.Printf("mqtt: can't connect to mqtt broker - %s\n", token.Error())
		logger.Error.Fatal("EXITING")
	}

	wg.Add(3)
	go subscribe(conn, data)
	go publish(conn, data.PubTopic+"volume", VolStatus)
	go publish(conn, data.PubTopic+"online", PcStatus)
	wg.Wait()
}

func publish(client mqtt.Client, topic string, f pubFunc) {
	i := 0
	for i < 1 {
		data, err := f()
		if err != nil {
			logger.Warn.Println("skiping one cycle of publishing")
			logger.Warn.Println(err)
		} else {
			token := client.Publish(topic, 0, false, data)
			token.Wait()
		}
		time.Sleep(StateInterval * time.Second)
	}

}

func subscribe(client mqtt.Client, data MqttConf) {
	topic := data.SubTopic + "#"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	logger.Info.Printf("subscribed to %q\n", topic)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	logger.Info.Println("connection to mqtt broker is successful")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logger.Error.Println("mqtt: connection to mqtt broker is lost")
	logger.Error.Fatal("EXITING")
}

func executor(topic, msg, subPrefix string) {
	logger.Info.Printf("%s recieved in %q\n", msg, topic)
	switch topic {
	case subPrefix + "volume":
		intMsg, err := strconv.Atoi(msg)
		if err == nil {
			sound.SetVolume(intMsg)
		} else {
			logger.Warn.Println("message must be in range of 0-100, skiping command")
		}
	case subPrefix + "mute":
		boolMsg, err := strconv.ParseBool(msg)
		if err == nil {
			sound.Mute(boolMsg)
		} else {
			logger.Warn.Println("message must be true or false, skiping command")
		}
	case subPrefix + "shutdown":
		power.Shutdown()
	case subPrefix + "reboot":
		power.Reboot()
	}
}
