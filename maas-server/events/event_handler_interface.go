package events

type EventHandlerInterface interface {
	//TODO should the username be passed in????
	ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32)
	//SORRY adding the following extension to the interface is an evil hack, abusing the concept
	// but required to implement the desired behavior with minimal change
	//If anybody feels like it, please feel free to improve!
	TotalItemConsumedForUserChanged(userID uint32, username string, itemID uint32, itemName string, itemCost int32, totalCount uint32)
	//SORRY adding the following extension to the interface is an even MORE evil hack, abusing the concept even harder
	// but required to implement the desired behavior with minimal change
	//If anybody feels like it, please feel free to improve!
	TotalItemConsumedChanged(itemID uint32, itemName string, itemCost int32, totalCount uint32)
}
