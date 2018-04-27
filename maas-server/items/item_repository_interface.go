package items

type ItemRepositoryInterface interface {
	Get(id uint32) (Item, error)
	List() ([]Item, error)
	Save(item Item) (Item, error)
	Delete(id uint32) (Item, error)
}
