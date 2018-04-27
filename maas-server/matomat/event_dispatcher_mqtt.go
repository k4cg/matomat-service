package matomat

import (
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type EventDispatcherMqtt struct {
	connectionString string
	clientID         string
	topic            string
	enabled          bool
	client           MQTT.Client
}

func NewEventDispatcherMqtt(connectionString string, clientID string, topic string, enabled bool) *EventDispatcherMqtt {
	opts := MQTT.NewClientOptions().AddBroker(connectionString) //connectionString example: "tcp://localhost:4242"
	opts.SetClientID(clientID)                                  //clientID example: "matomat-server"
	return &EventDispatcherMqtt{connectionString: connectionString, clientID: clientID, topic: topic, client: MQTT.NewClient(opts), enabled: enabled}
}

//TODO should the username be passed in????
func (ed *EventDispatcherMqtt) ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost uint32) error {
	var err error
	if !ed.enabled {
		return err
	}

	//start a mqtt client
	if token := ed.client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	} else {
		token := ed.client.Publish(ed.topic, 0, false, buildItemConsumedMessage(userID, username, itemID, itemName, itemCost))
		//wait for receipt from broker
		token.Wait()
		ed.client.Disconnect(250)
	}

	return err
}

func buildItemConsumedMessage(userID uint32, username string, itemID uint32, itemName string, itemCost uint32) string {
	message := "matomat;item-consumed;" + strconv.Itoa(int(userID)) + ";" + strconv.Itoa(int(itemID)) + ";" + itemName + ";" + strconv.Itoa(int(itemCost)) //TODO implement proper message format
	return message
}
