package events

type ItemConsumedMessage struct {
	EventSource string
	EventName   string
	EventTime   int64
	UserID      uint32
	ItemID      uint32
	ItemName    string
	ItemCost    int32
	ItemCount   uint32
}

type TotalItemConsumedForUserChangedMessage struct {
	EventSource    string
	EventName      string
	EventTime      int64
	UserID         uint32
	ItemID         uint32
	ItemName       string
	ItemCost       int32
	TotalItemCount uint32
}

type TotalItemConsumedChangedMessage struct {
	EventSource    string
	EventName      string
	EventTime      int64
	ItemID         uint32
	ItemName       string
	ItemCost       int32
	TotalItemCount uint32
}

func NewItemConsumedMessage(eventTime int64, userID uint32, itemID uint32, itemName string, itemCost int32, count uint32) *ItemConsumedMessage {
	return &ItemConsumedMessage{EventSource: "matomat", EventName: "item-consumed", EventTime: eventTime, UserID: userID, ItemID: itemID, ItemName: itemName, ItemCost: itemCost, ItemCount: count}
}

func NewTotalItemConsumedForUserChangedMessage(eventTime int64, userID uint32, itemID uint32, itemName string, itemCost int32, totalCount uint32) *TotalItemConsumedForUserChangedMessage {
	return &TotalItemConsumedForUserChangedMessage{EventSource: "matomat", EventName: "total-item-consumed-for-user-changed", EventTime: eventTime, UserID: userID, ItemID: itemID, ItemName: itemName, ItemCost: itemCost, TotalItemCount: totalCount}
}

func NewTotalItemConsumedChangedMessage(eventTime int64, itemID uint32, itemName string, itemCost int32, totalCount uint32) *TotalItemConsumedChangedMessage {
	return &TotalItemConsumedChangedMessage{EventSource: "matomat", EventName: "total-items-consumed", EventTime: eventTime, ItemID: itemID, ItemName: itemName, ItemCost: itemCost, TotalItemCount: totalCount}
}
