package users

import (
	"database/sql"

	"github.com/k4cg/matomat-service/maas-server/items"
	_ "github.com/mattn/go-sqlite3"
)

type UserItemsStatsRepoSqlite3 struct {
	db *sql.DB
}

func NewUserItemsStatsRepoSqlite3(sqlite3DbFilePath string) *UserItemsStatsRepoSqlite3 {
	db, err := sql.Open("sqlite3", sqlite3DbFilePath)
	if err == nil {
		return &UserItemsStatsRepoSqlite3{db: db}
	} else {
		panic(err)
	}
}

func (r *UserItemsStatsRepoSqlite3) Get(userID uint32) (map[uint32]items.ItemStats, error) {
	itemsStats := make(map[uint32]items.ItemStats)
	var err error

	rows, err := r.db.Query("SELECT itemID, consumed FROM user_items_stats WHERE userID=?")
	if err == nil {
		for rows.Next() {
			var itemid uint32
			var consumed uint32

			rows.Scan(&itemid, &consumed)

			itemsStats[itemid] = items.ItemStats{ItemID: itemid, Consumed: consumed}
		}
		rows.Close()
	}

	return itemsStats, err
}

func (r *UserItemsStatsRepoSqlite3) CountConsumption(userID uint32, itemID uint32, consumed uint32) error {
	var err error
	oldItemStats, err := r.Get(userID)
	if err == nil {
		_, found := oldItemStats[itemID]
		if !found {
			//create new stats entry
			stmt, err := r.db.Prepare("INSERT INTO user_items_stats (userID, itemID, consumed) VALUES (?, ?, ?)")
			if err == nil {
				_, err = stmt.Exec(userID, itemID, consumed)
			}
			stmt.Close()
		} else {
			stmt, err := r.db.Prepare("UPDATE user_items_stats SET consumed = consumed + ? WHERE userID=? AND itemID=?")
			if err == nil {
				_, err = stmt.Exec(consumed, userID, itemID)
			}
			stmt.Close()
		}
	}
	return err
}
