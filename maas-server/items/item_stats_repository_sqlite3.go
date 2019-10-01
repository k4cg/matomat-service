package items

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ItemStatsRepoSqlite3 struct {
	db *sql.DB
}

func NewItemStatsRepoSqlite3(sqlite3DbFilePath string) *ItemStatsRepoSqlite3 {
	db, err := sql.Open("sqlite3", sqlite3DbFilePath)
	if err == nil {
		return &ItemStatsRepoSqlite3{db: db}
	} else {
		panic(err)
	}
}

func (r *ItemStatsRepoSqlite3) Get(itemID uint32) (ItemStats, error) {
	var itemStats ItemStats
	var err error

	rows, err := r.db.Query("SELECT itemID, consumed FROM items_stats WHERE itemID=?", itemID)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&itemStats.ItemID, &itemStats.Consumed)
			break
		}
		err = rows.Close()
	}

	return itemStats, err
}

func (r *ItemStatsRepoSqlite3) CountConsumption(itemID uint32, consumed uint32) (ItemStats, error) {
	var returnedItemStats ItemStats
	var err error
	oldItemStats, err := r.Get(itemID)
	if err == nil {
		if oldItemStats == (ItemStats{}) {
			//create new stats entry
			stmt, err := r.db.Prepare("INSERT INTO items_stats (itemID, consumed) VALUES (?, ?)")
			if err == nil {
				_, err := stmt.Exec(itemID, consumed)
				if err == nil {
					returnedItemStats = ItemStats{ItemID: itemID, Consumed: consumed}
				}
				err = stmt.Close()
			}
		} else {
			stmt, err := r.db.Prepare("UPDATE items_stats SET consumed = consumed + ? WHERE itemID=?")
			if err == nil {
				_, err = stmt.Exec(consumed, itemID)
				if err == nil {
					returnedItemStats = ItemStats{ItemID: itemID, Consumed: consumed}
				}
				err = stmt.Close()
			}
		}
	}
	return returnedItemStats, err
}

func (r *ItemStatsRepoSqlite3) List() ([]ItemStats, error) {
	itemsStats := make([]ItemStats, 0)
	var err error

	rows, err := r.db.Query("SELECT itemID, consumed FROM items_stats")
	if err == nil {
		for rows.Next() {
			var id uint32
			var consumed uint32

			err = rows.Scan(&id, &consumed)

			if err == nil {
				itemsStats = append(itemsStats, ItemStats{ItemID: id, Consumed: consumed})
			}
		}
		err = rows.Close()
	}

	return itemsStats, err
}
