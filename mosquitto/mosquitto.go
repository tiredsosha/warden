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
	keyLifeTime   = 2  // minute
	reconnTime    = 20 // sec
	pubVolumeTime = 5  // sec
	pubMuteTime   = 7  // sec
)

type pubFunc func() (string, error)

type MqttConf struct {
	ID       string
	Broker   string
	Username string
	Password string
	SubTopic string
	PubTopic string
}

func (data *MqttConf) messageHandler(client mqtt.Client, msgHand mqtt.Message) {
	topic := msgHand.Topic()
	msg := strings.TrimSpace(string(msgHand.Payload()))
	executorer(topic, msg, data.SubTopic)
}

func (data *MqttConf) connectHandler(client mqtt.Client) {
	topic := data.SubTopic + "#"
	client.Unsubscribe(topic)

	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	logger.Debug.Printf("publishing to %q\n", data.PubTopic+"#")
	logger.Debug.Printf("subscribed to %q\n", topic)
	logger.Info.Println("connection to mqtt broker is successful")
	tokenPub := client.Publish(data.PubTopic+"online", 0, true, "true")
	tokenPub.Wait()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logger.Warn.Printf("mqtt: connection to mqtt broker is lost - %s\n", err)
}

func StartBroker(data MqttConf) {
	var wg sync.WaitGroup

	messagePubHandler := data.messageHandler
	connectHandler := data.connectHandler

	// MQTT INIT //
	mqttHandler := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", data.Broker, port)).
		SetClientID(data.ID).
		SetUsername(data.Username).
		SetPassword(data.Password).
		SetAutoReconnect(true).
		SetDefaultPublishHandler(messagePubHandler).
		SetConnectionLostHandler(connectLostHandler).
		SetOnConnectHandler(connectHandler).
		SetKeepAlive(keyLifeTime).
		SetWill(data.PubTopic+"online", "false", 2, true)

	conn := mqtt.NewClient(mqttHandler)

	for {
		status := true
		if token := conn.Connect(); token.Wait() && token.Error() != nil {
			logger.Warn.Printf("mqtt: can't connect to mqtt broker - %s\n", token.Error())
			status = false
		}

		if status {
			break
		}
		time.Sleep(reconnTime * time.Second)
	}

	// ADD YOUR PUBLISHERS //
	wg.Add(2)
	go publisher(conn, data.PubTopic+"volume", VolumeStatus, pubVolumeTime)
	go publisher(conn, data.PubTopic+"muted", MuteStatus, pubMuteTime)
	wg.Wait()
}

func publisher(client mqtt.Client, topic string, f pubFunc, sleep int) {
	for {
		data, err := f()
		if err != nil {
			logger.Warn.Printf("skiping one cycle of publishing to %q - %s\n", topic, err)
		} else {
			token := client.Publish(topic, 0, false, data)
			token.Wait()
		}
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

// HANDLER FOR COMMAND IN DIFFERENT TOPICS //
func executorer(topic, msg, subPrefix string) {
	logger.Debug.Printf("%s recieved in %q\n", msg, topic)
	switch topic {
	case subPrefix + "volume":
		intMsg, err := strconv.Atoi(msg)
		if err == nil {
			sound.SetVolume(intMsg)
		} else {
			logger.Warn.Println("message in volume topic must be in range of 0-100, skiping command")
		}
	case subPrefix + "mute":
		boolMsg, err := strconv.ParseBool(msg)
		if err == nil {
			sound.Mute(boolMsg)
		} else {
			logger.Warn.Println("message in mute topic must be true or false, skiping command")
		}
	case subPrefix + "shutdown":
		power.Shutdown()
	case subPrefix + "reboot":
		power.Reboot()
		// case subPrefix + "display":
		// 	boolMsg, err := strconv.ParseBool(msg)
		// 	if err == nil {
		// 		power.Display(boolMsg)
		// 	} else {
		// 		logger.Warn.Println("message in mute topic must be true or false, skiping command")
		// 	}
	}
}
