package events

import (
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/users"
)

type EventHandlerStatsRepos struct {
	itemStatsRepo      items.ItemStatsRepositoryInterface
	userItemsStatsRepo users.UserItemsStatsRepositoryInterface
	//SORRY adding the following is an evil hack, abusing the concept
	// but required to implement the desired behavior with minimal change
	//If anybody feels like it, please feel free to improve!
	eventHandlerMessaging EventHandlerInterface
}

func NewEventHandlerStatsRepos(itemStatsRepo items.ItemStatsRepositoryInterface, userItemsStatsRepo users.UserItemsStatsRepositoryInterface, eventHandlerMessaging EventHandlerInterface) *EventHandlerStatsRepos {
	return &EventHandlerStatsRepos{itemStatsRepo: itemStatsRepo, userItemsStatsRepo: userItemsStatsRepo, eventHandlerMessaging: eventHandlerMessaging}
}

func (eh *EventHandlerStatsRepos) ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) {
	_, _ = eh.itemStatsRepo.CountConsumption(itemID, count)
	_ = eh.userItemsStatsRepo.CountConsumption(userID, itemID, count)

	//SORRY adding the following is an evil hack, abusing the concept
	// but required to implement the desired behavior with minimal change
	//If anybody feels like it, please feel free to improve!
	userItemStatsList, _ := eh.userItemsStatsRepo.Get(userID)

	var totalCount uint32 = 0
	for _, userItemStats := range userItemStatsList {
		if userItemStats.ItemID == itemID {
			totalCount = userItemStats.Consumed
			break
		}
	}

	eh.eventHandlerMessaging.TotalItemConsumedForUserChanged(userID, username, itemID, itemName, itemCost, totalCount)
}

//SORRY adding the following is an evil hack, abusing the concept
// but required to implement the desired behavior with minimal change
//If anybody feels like it, please feel free to improve!
func (eh *EventHandlerStatsRepos) TotalItemConsumedForUserChanged(userID uint32, username string, itemID uint32, itemName string, itemCost int32, totalCount uint32) {
	//do nothing ... as said, evil hack, badly implemented but will work
}
