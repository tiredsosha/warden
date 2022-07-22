package mosquitto

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tiredsosha/wardener/control/power"
	"github.com/tiredsosha/wardener/control/sound"
)

const (
	port          = 1883
	KeyLifeTime   = 2  // minute
	StateInterval = 10 // second
)

type MqttConf struct {
	Id       string
	Broker   string
	Username string
	Password string
	SubTopic string
	PubTopic string
}

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

	subscribe(conn, data)
	publish(conn, data)
	conn.Disconnect(250)
}

func publish(client mqtt.Client, data MqttConf) {
	i := 0
	for i < 1 {
		// volume, err := sound.GetVolume()
		// if err == nil {
		volume, err := sound.GetVolume()
		fmt.Println(reflect.TypeOf(volume), err)

		token := client.Publish(data.PubTopic, 0, false, "1")
		token.Wait()
		time.Sleep(time.Second)
		// token := client.Publish(data.PubTopic, 0, false, volume)
		// token.Wait()
		// time.Sleep(StateInterval)
		//}
	}
}

func subscribe(client mqtt.Client, data MqttConf) {
	token := client.Subscribe(data.SubTopic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", data.SubTopic)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msgHand mqtt.Message) {
	topic := msgHand.Topic()
	msg := string(msgHand.Payload())
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
	case "wardener/sound/volume":
		intMsg, err := strconv.Atoi(msg)
		if err == nil {
			sound.SetVolume(intMsg)
		}
	case "wardener/sound/mute":
		boolMsg, err := strconv.ParseBool(msg)
		if err == nil {
			sound.Mute(boolMsg)
		}
	case "wardener/power/shutdown":
		power.Shutdown()
	case "wardener/power/reboot":
		power.Reboot()
	default:
		fmt.Printf("%s recieved in %d\n", msg, topic)
	}
}
