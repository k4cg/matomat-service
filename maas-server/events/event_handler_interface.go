package events

type EventHandlerInterface interface {
	//TODO add timestamp?
	//TODO should the username be passed in????
	ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32)
}
