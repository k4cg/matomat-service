package items

type ItemStatsRepositoryInterface interface {
	Get(id uint32) (ItemStats, error)
	List() (map[uint32]ItemStats, error)
	CountConsumption(itemID uint32, consumed uint32) (ItemStats, error)
}
