package items

import (
	"errors"
)

type ItemRepoMem struct {
	autoIncrement uint32
	items         map[uint32]Item
	itemsDeleted  map[uint32]Item
}

func NewItemRepoMem() *ItemRepoMem {
	return &ItemRepoMem{items: make(map[uint32]Item), itemsDeleted: make(map[uint32]Item), autoIncrement: 1}
}

func (r *ItemRepoMem) Get(id uint32) (Item, error) {
	var err error
	item, found := r.items[id]
	if !found {
		err = errors.New("No item found with given item ID")
	}
	return item, err
}

func (r *ItemRepoMem) List() (map[uint32]Item, error) {
	return r.items, nil
}

func (r *ItemRepoMem) Save(item Item) (Item, error) {
	var err error
	if item.ID != 0 {
		r.items[item.ID] = item
	} else {
		item.ID = r.autoIncrement
		r.autoIncrement = r.autoIncrement + 1
		r.items[item.ID] = item
	}
	return item, err
}

func (r *ItemRepoMem) Delete(id uint32) (Item, error) {
	item, found := r.items[id]
	if found {
		delete(r.items, id)
		r.itemsDeleted[id] = item
	}
	return item, nil
}
