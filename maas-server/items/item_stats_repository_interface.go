package items

type ItemStatsRepositoryInterface interface {
	Get(id uint32) (ItemStats, error)
	List() ([]ItemStats, error)
	CountConsumption(itemID uint32, consumed uint32) (ItemStats, error)
}
