package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/karloygard/xcomfortd-go/xc"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttRelay struct {
	xc.Interface

	client mqtt.Client
	ctx    context.Context
}

func (r *MqttRelay) dimmerCallback(c mqtt.Client, msg mqtt.Message) {
	var dp, value int

	if _, err := fmt.Sscanf(msg.Topic(), "xcomfort/%d/set/dimmer", &dp); err != nil {
		log.Println(err)
		return
	}
	if _, err := fmt.Sscanf(string(msg.Payload()), "%d", &value); err != nil {
		log.Println(err)
		return
	}

	if datapoint := r.Datapoint(dp); datapoint != nil {
		log.Printf("topic: %s, message: %s\n", msg.Topic(), string(msg.Payload()))

		if _, err := datapoint.Dim(r.ctx, value); err != nil {
			log.Println(err)
		} else {
			r.StatusValue(dp, value)
			// Send bool as well, to appease HA
			r.StatusBool(dp, value > 0)
		}
	} else {
		log.Printf("unknown datapoint %d\n", dp)
	}
}

func (r *MqttRelay) switchCallback(c mqtt.Client, msg mqtt.Message) {
	var dp int

	if _, err := fmt.Sscanf(msg.Topic(), "xcomfort/%d/set/switch", &dp); err != nil {
		log.Println(err)
		return
	}

	if datapoint := r.Datapoint(dp); datapoint != nil {
		log.Printf("topic: %s, message: %s\n", msg.Topic(), string(msg.Payload()))

		on := string(msg.Payload()) == "true"

		if _, err := datapoint.Switch(r.ctx, on); err != nil {
			log.Println(err)
		} else {
			r.StatusBool(dp, on)
		}
	} else {
		log.Printf("unknown datapoint %d\n", dp)
	}
}

func (r *MqttRelay) StatusValue(datapoint, value int) {
	topic := fmt.Sprintf("xcomfort/%d/get/dimmer", datapoint)
	r.client.Publish(topic, 1, true, fmt.Sprint(value))
	r.StatusBool(datapoint, value > 0)
}

func (r *MqttRelay) StatusBool(datapoint int, on bool) {
	topic := fmt.Sprintf("xcomfort/%d/get/switch", datapoint)
	r.client.Publish(topic, 1, true, fmt.Sprint(on))
}

func (r *MqttRelay) Connect(ctx context.Context, clientId string, uri *url.URL) error {
	opts := mqtt.NewClientOptions()
	broker := fmt.Sprintf("tcp://%s", uri.Host)

	log.Printf("using broker %s", broker)

	opts.AddBroker(broker)
	opts.SetUsername(uri.User.Username())
	if password, set := uri.User.Password(); set {
		opts.SetPassword(password)
	}
	opts.SetClientID(clientId)

	r.client = mqtt.NewClient(opts)
	token := r.client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return err
	}

	r.client.Subscribe("xcomfort/+/set/dimmer", 0, r.dimmerCallback)
	r.client.Subscribe("xcomfort/+/set/switch", 0, r.switchCallback)

	r.ctx = ctx

	return nil
}