package items

import (
	"database/sql"
	"errors"
)

type ItemRepoSqlite3 struct {
	db *sql.DB
}

func NewItemRepoSqlite3(sqlite3DbFilePath string) *ItemRepoSqlite3 {
	db, err := sql.Open("sqlite3", sqlite3DbFilePath)
	if err == nil {
		return &ItemRepoSqlite3{db: db}
	} else {
		panic(err)
	}
}

func (r *ItemRepoSqlite3) Get(itemID uint32) (Item, error) {
	var item Item
	var err error

	rows, err := r.db.Query("SELECT ID, name, cost FROM items WHERE id=?", itemID)
	if err == nil {
		for rows.Next() {
			rows.Scan(&item.ID, &item.Name, &item.Cost)
			break
		}
	}

	if item == (Item{}) {
		err = errors.New("No item found for given item ID")
	}

	return item, err
}

func (r *ItemRepoSqlite3) List() (map[uint32]Item, error) {
	var items map[uint32]Item
	var err error

	rows, err := r.db.Query("SELECT ID, name, cost FROM items")
	if err == nil {
		for rows.Next() {
			var id uint32
			var name string
			var cost uint32
			rows.Scan(&id, &name, &cost)
			items[id] = Item{ID: id, Name: name, Cost: cost}
		}
	}

	return items, err
}

func (r *ItemRepoSqlite3) Save(item Item) (Item, error) {
	var returnedItem Item
	var err error
	oldItem, err := r.Get(item.ID)
	if err == nil {
		if oldItem == (Item{}) {
			//create new item
			stmt, err := r.db.Prepare("INSERT INTO items (name, cost) VALUES (?, ?)")
			if err == nil {
				res, err := stmt.Exec(item.Name, item.Cost)
				id, err := res.LastInsertId()
				if err == nil {
					//evil cast of int64 => uint32 .... TODO solve this better
					returnedItem = Item{ID: uint32(id), Name: item.Name, Cost: item.Cost}
				}
			}
		} else {
			stmt, err := r.db.Prepare("UPDATE items SET name=?, cost=? WHERE ID=?")
			if err == nil {
				_, err = stmt.Exec(item.Name, item.Cost, item.ID)
			}
		}
	}
	return returnedItem, err
}

func (r *ItemRepoSqlite3) Delete(itemID uint32) (Item, error) {
	var item Item
	var err error
	item, err = r.Get(itemID)
	if err == nil {
		stmt, err := r.db.Prepare("DELETE FROM items WHERE ID=? LIMIT 1")
		if err == nil {
			_, err = stmt.Exec(item.ID)
		}
	}
	return item, err
}
