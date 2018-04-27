package users

import "github.com/k4cg/matomat-service/maas-server/items"

type UserItemsStatsRepositoryInterface interface {
	Get(userID uint32) ([]items.ItemStats, error)
	CountConsumption(userID uint32, itemID uint32, consumed uint32) error
}
