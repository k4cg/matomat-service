package items

func GetStatsForItem(itemStatsList []ItemStats, itemID uint32) (ItemStats, bool) {
	var returnedItemStats ItemStats
	found := false

	for _, itemStats := range itemStatsList {
		if itemStats.ItemID == itemID {
			returnedItemStats = itemStats
			found = true
			break
		}
	}

	return returnedItemStats, found
}
