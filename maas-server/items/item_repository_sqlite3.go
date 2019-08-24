package items

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
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

	rows, err := r.db.Query("SELECT ID, name, cost, stock FROM items WHERE id=?", itemID)
	if err == nil {
		for rows.Next() {
			rows.Scan(&item.ID, &item.Name, &item.Cost, &item.Stock)
			break
		}
		rows.Close()
	}

	return item, err
}

func (r *ItemRepoSqlite3) List() ([]Item, error) {
	items := make([]Item, 0)
	var err error

	rows, err := r.db.Query("SELECT ID, name, cost, stock FROM items")
	if err == nil {
		for rows.Next() {
			var id uint32
			var name string
			var cost int32
			var stock int32
			rows.Scan(&id, &name, &cost, &stock)
			items = append(items, Item{ID: id, Name: name, Cost: cost, Stock: stock})
		}
		rows.Close()
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
			stmt, err := r.db.Prepare("INSERT INTO items (name, cost, stock) VALUES (?, ?, ?)")
			if err == nil {
				res, err := stmt.Exec(item.Name, item.Cost, item.Stock)
				id, err := res.LastInsertId()
				if err == nil {
					//evil cast of int64 => uint32 .... TODO solve this better
					returnedItem = Item{ID: uint32(id), Name: item.Name, Cost: item.Cost, Stock: item.Stock}
				}
				stmt.Close()
			}
		} else {
			stmt, err := r.db.Prepare("UPDATE items SET name=?, cost=?, stock=? WHERE ID=?")
			if err == nil {
				_, err = stmt.Exec(item.Name, item.Cost, item.Stock, item.ID)
				if err == nil {
					returnedItem = item
				}
				stmt.Close()
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
		stmt, err := r.db.Prepare("DELETE FROM items WHERE ID=?")
		if err == nil {
			_, err = stmt.Exec(item.ID)
			stmt.Close()
		}
	}
	return item, err
}
