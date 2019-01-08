package matomat

import (
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/users"
)

type ServiceStats struct {
	Items *ItemsServiceStats
	Users *UsersServiceStats
}

type ItemsServiceStats struct {
	Count uint32
	Cost  *ItemsCostServiceStats
}

type ItemsCostServiceStats struct {
	Avg int32
	Min int32
	Max int32
}

type UsersServiceStats struct {
	Count   uint32
	Credits *UsersCreditsServiceStats
}

type UsersCreditsServiceStats struct {
	Sum int32
	Avg int32
	Min int32
	Max int32
}

func (m *Matomat) ServiceStatsGet() (*ServiceStats, error) {
	var err error

	items, _ := m.itemRepo.List()
	users, _ := m.userRepo.List()

	itemsServiceStats := &ItemsServiceStats{Count: uint32(len(items)), Cost: m.getItemsCostServiceStats(items)}
	usersServiceStats := &UsersServiceStats{Count: uint32(len(users)), Credits: m.getUsersCreditsServiceStats(users)}

	stats := &ServiceStats{Items: itemsServiceStats, Users: usersServiceStats}

	return stats, err
}

func (m *Matomat) getItemsCostServiceStats(items []items.Item) *ItemsCostServiceStats {
	var avg int32
	sum := 0
	min := ^int32(0)
	max := int32(0)

	for _, item := range items {
		sum += int(item.Cost)
		if item.Cost > max {
			max = item.Cost
		}
		if item.Cost < min {
			min = item.Cost
		}
	}

	avg = 0
	if len(items) > 0 {
		avg = int32(sum / len(items)) //another evil cast AND this is cut off, so average is only credit unit exact
	}

	return &ItemsCostServiceStats{Avg: avg, Min: min, Max: max}
}

func (m *Matomat) getUsersCreditsServiceStats(users []users.User) *UsersCreditsServiceStats {
	var avg int32
	sum := 0
	min := ^int32(0)
	max := int32(0)

	for _, user := range users {
		sum += int(user.Credits)
		if user.Credits > max {
			max = user.Credits
		}
		if user.Credits < min {
			min = user.Credits
		}
	}

	avg = 0
	if len(users) > 0 {
		avg = int32(sum / len(users)) //another evil cast AND this is cut off, so average is only credit unit exact
	}

	return &UsersCreditsServiceStats{Sum: int32(sum), Avg: avg, Min: min, Max: max}
}
