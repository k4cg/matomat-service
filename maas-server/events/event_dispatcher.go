package events

type EventDispatcher struct {
	handlers []EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{}
}

func (ed *EventDispatcher) Register(handler EventHandlerInterface) {
	ed.handlers = append(ed.handlers, handler)
}

//TODO should the username be passed in????
func (ed *EventDispatcher) ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) {
	for _, handler := range ed.handlers {
		go handler.ItemConsumed(userID, username, itemID, itemName, itemCost, count)
	}
}
