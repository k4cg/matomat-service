package items

import (
	"database/sql"
	"errors"
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
			rows.Scan(&itemStats.ItemID, &itemStats.Consumed)
			break
		}
	}

	if itemStats == (ItemStats{}) {
		err = errors.New("No item stats found for given item ID")
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
			}
		} else {
			stmt, err := r.db.Prepare("UPDATE items_stats SET consumed = consumed + ? WHERE itemID=?")
			if err == nil {
				_, err = stmt.Exec(consumed, itemID)
				if err == nil {
					returnedItemStats = ItemStats{ItemID: itemID, Consumed: consumed}
				}
			}
		}
	}
	return returnedItemStats, err
}

func (r *ItemStatsRepoSqlite3) List() (map[uint32]ItemStats, error) {
	var itemsStats map[uint32]ItemStats
	var err error

	rows, err := r.db.Query("SELECT itemID, consumed FROM items_stats")
	if err == nil {
		for rows.Next() {
			var id uint32
			var consumed uint32

			rows.Scan(&id, &consumed)

			itemsStats[id] = ItemStats{ItemID: id, Consumed: consumed}
		}
	}

	return itemsStats, err
}
