package events

import (
	"log"
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type EventHandlerMqtt struct {
	connectionString string
	clientID         string
	topic            string
	enabled          bool
	clientOpts       *MQTT.ClientOptions
}

func NewEventHandlerMqtt(connectionString string, clientID string, topic string, enabled bool, username string, password string) *EventHandlerMqtt {
	clientOpts := MQTT.NewClientOptions().AddBroker(connectionString) //connectionString example: "tcp://localhost:4242"
	clientOpts.SetClientID(clientID)                                  //clientID example: "matomat-server"
	if len(username) > 0 {
		clientOpts.SetUsername(username)
		clientOpts.SetPassword(password)
	}
	return &EventHandlerMqtt{connectionString: connectionString, clientID: clientID, topic: topic, enabled: enabled, clientOpts: clientOpts}
}

func (eh *EventHandlerMqtt) ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) {
	client := MQTT.NewClient(eh.clientOpts)
	// connect to the broker using the mqtt client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err := token.Error()
		if err != nil {
			log.Print("EventHandlerMqtt: error trying to connect to MQTT broker")
		}
	} else {
		token := client.Publish(eh.topic, 0, false, eh.buildItemConsumedMessage(userID, username, itemID, itemName, itemCost, count))
		//wait for receipt from broker
		token.Wait()
		client.Disconnect(250)
	}
}

func (eh *EventHandlerMqtt) buildItemConsumedMessage(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) string {
	//TODO - shouldn't the username used instead of the user ID?
	//TODO - the int casts are messy ... improve!
	message := "matomat;item-consumed;" + strconv.Itoa(int(userID)) + ";" + strconv.Itoa(int(itemID)) + ";" + itemName + ";" + strconv.Itoa(int(itemCost)) + ";" + strconv.Itoa(int(count)) //TODO implement proper message format
	return message
}
