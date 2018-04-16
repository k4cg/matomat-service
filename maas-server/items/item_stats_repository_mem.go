package items

import (
	"errors"
)

type ItemStatsRepoMem struct {
	autoIncrement uint32
	itemStats     map[uint32]ItemStats
}

func NewItemStatsRepoMem() *ItemStatsRepoMem {
	return &ItemStatsRepoMem{itemStats: make(map[uint32]ItemStats)}
}

func (r *ItemStatsRepoMem) Get(itemID uint32) (ItemStats, error) {
	var err error
	itemStats, found := r.itemStats[itemID]
	if !found {
		err = errors.New("No item stats found for given item ID")
	}
	return itemStats, err
}

func (r *ItemStatsRepoMem) CountConsumption(itemID uint32, consumed uint32) (ItemStats, error) {
	var err error
	itemStats, found := r.itemStats[itemID]
	if found {
		itemStats.Consumed = itemStats.Consumed + consumed
	} else {
		itemStats = ItemStats{ItemID: itemID, Consumed: consumed}
		r.itemStats[itemID] = itemStats
	}
	return itemStats, err
}

func (r *ItemStatsRepoMem) List() (map[uint32]ItemStats, error) {
	return r.itemStats, nil
}
