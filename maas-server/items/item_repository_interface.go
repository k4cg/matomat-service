package items

type ItemRepositoryInterface interface {
	Get(id uint32) (Item, error)
	List() (map[uint32]Item, error)
	Save(item Item) (Item, error)
	Delete(id uint32) (Item, error)
}
