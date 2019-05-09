package events

import (
	"log"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type EventHandlerMqtt struct {
	connectionString string
	clientID         string
	topic            string
	clientOpts       *MQTT.ClientOptions
	retainMessage    bool
}

//TODO - switch constructor to using config struct instead of droelf parameters
func NewEventHandlerMqtt(connectionString string, clientID string, topic string, username string, password string, retainMessage bool) *EventHandlerMqtt {
	clientOpts := MQTT.NewClientOptions().AddBroker(connectionString) //connectionString example: "tcp://localhost:4242"
	clientOpts.SetClientID(clientID)                                  //clientID example: "matomat-server"
	if len(username) > 0 {
		clientOpts.SetUsername(username)
		clientOpts.SetPassword(password)
	}
	return &EventHandlerMqtt{connectionString: connectionString, clientID: clientID, topic: topic, clientOpts: clientOpts, retainMessage: retainMessage}
}

func (eh *EventHandlerMqtt) ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) {
	eh.publishMessage(eh.buildItemConsumedMessage(userID, username, itemID, itemName, itemCost, count))
}

func (eh *EventHandlerMqtt) buildItemConsumedMessage(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) string {
	//TODO - shouldn't the username used instead of the user ID?
	//TODO - the int casts are messy ... improve!
	message := "matomat;item-consumed;" + strconv.Itoa(int(userID)) + ";" + strconv.Itoa(int(itemID)) + ";" + itemName + ";" + strconv.Itoa(int(itemCost)) + ";" + strconv.Itoa(int(count)) + ";" + strconv.FormatInt(time.Now().Unix(), 10) //TODO implement proper message format
	return message
}

//SORRY adding the following is an evil hack, abusing the concept
// but required to implement the desired behavior with minimal change
//If anybody feels like it, please feel free to improve!
func (eh *EventHandlerMqtt) TotalItemConsumedForUserChanged(userID uint32, username string, itemID uint32, itemName string, itemCost int32, totalCount uint32) {
	eh.publishMessage(eh.buildTotalItemConsumedForUserChanged(userID, username, itemID, itemName, itemCost, totalCount))
}

func (eh *EventHandlerMqtt) buildTotalItemConsumedForUserChanged(userID uint32, username string, itemID uint32, itemName string, itemCost int32, totalCount uint32) string {
	//TODO - shouldn't the username used instead of the user ID?
	//TODO - the int casts are messy ... improve!
	message := "matomat;total-item-consumed-for-user-changed;" + strconv.Itoa(int(userID)) + ";" + strconv.Itoa(int(itemID)) + ";" + itemName + ";" + strconv.Itoa(int(itemCost)) + ";" + strconv.Itoa(int(totalCount)) + ";" + strconv.FormatInt(time.Now().Unix(), 10) //TODO implement proper message format
	return message
}

//SORRY adding the following is an even more evil hack, abusing the concept even more
// but required to implement the desired behavior with minimal change
//If anybody feels like it, please feel free to improve!
func (eh *EventHandlerMqtt) TotalItemConsumedChanged(itemID uint32, itemName string, itemCost int32, totalCount uint32) {
	eh.publishMessage(eh.buildTotalItemConsumedChanged(itemID, itemName, itemCost, totalCount))
}

func (eh *EventHandlerMqtt) buildTotalItemConsumedChanged(itemID uint32, itemName string, itemCost int32, totalCount uint32) string {
	message := "matomat;total-items-consumed;" + strconv.Itoa(int(itemID)) + ";" + itemName + ";" + strconv.Itoa(int(itemCost)) + ";" + strconv.Itoa(int(totalCount)) + ";" + strconv.FormatInt(time.Now().Unix(), 10) //TODO implement proper message format
	return message
}

func (eh *EventHandlerMqtt) publishMessage(message string) {
	client := MQTT.NewClient(eh.clientOpts)
	// connect to the broker using the mqtt client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err := token.Error()
		if err != nil {
			log.Print("EventHandlerMqtt: error trying to connect to MQTT broker")
		}
	} else {
		token := client.Publish(eh.topic, 0, eh.retainMessage, message)
		//wait for receipt from broker
		token.Wait()
		client.Disconnect(250)
	}
}
