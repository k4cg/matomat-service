package events

import (
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/users"
)

type EventHandlerStatsRepos struct {
	itemStatsRepo      items.ItemStatsRepositoryInterface
	userItemsStatsRepo users.UserItemsStatsRepositoryInterface
}

func NewEventHandlerStatsRepos(itemStatsRepo items.ItemStatsRepositoryInterface, userItemsStatsRepo users.UserItemsStatsRepositoryInterface) *EventHandlerStatsRepos {
	return &EventHandlerStatsRepos{itemStatsRepo: itemStatsRepo, userItemsStatsRepo: userItemsStatsRepo}
}

func (eh *EventHandlerStatsRepos) ItemConsumed(userID uint32, username string, itemID uint32, itemName string, itemCost int32, count uint32) {
	_, _ = eh.itemStatsRepo.CountConsumption(userID, count)
	_ = eh.userItemsStatsRepo.CountConsumption(userID, itemID, count)
}
