package api

import (
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/matomat"
	"github.com/k4cg/matomat-service/maas-server/users"
)

type Error struct {
	Message string `json:"message"`
}

func newUser(user users.User) User {
	return User{ID: user.ID, Username: user.Username, Credits: user.Credits, Admin: user.IsAdmin()}
}

//Represenation of a Matomat users on API level
type User struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Credits  int32  `json:"credits"`
	Admin    bool   `json:"admin"`
}

type TransferredCredits struct {
	Sender  User  `json:"sender"`
	Credits int32 `json:"credits"`
}

func newTransferredCredits(fromUser users.User, transferredCreditsAmount int32) TransferredCredits {
	return TransferredCredits{Sender: newUser(fromUser), Credits: transferredCreditsAmount}
}

//Item Represenation of a drink/pizza/thing that is being "sold" by Matomat, containing its name and cost (in credits).
type Item struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Cost int32  `json:"cost"`
}

func newItem(item items.Item) Item {
	return Item{ID: item.ID, Name: item.Name, Cost: item.Cost}
}

type ItemStats struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Cost     int32  `json:"cost"`
	Consumed uint32 `json:"consumed"`
}

func newItemStats(item items.Item, itemStats items.ItemStats) ItemStats {
	return ItemStats{ID: item.ID, Name: item.Name, Cost: item.Cost, Consumed: itemStats.Consumed}
}

type AuthSuccess struct {
	Token               string `json:"token"`
	ExpirationTimestamp uint32 `json:"expires"`
	User                User   `json:"user"`
}

func newAuthSuccess(token string, expirationTimestamp uint32, user users.User) AuthSuccess {
	return AuthSuccess{Token: token, ExpirationTimestamp: expirationTimestamp, User: newUser(user)}
}

type ServiceStats struct {
	Items *ItemsServiceStats `json:"items"`
	Users *UsersServiceStats `json:"users"`
}

type ItemsServiceStats struct {
	Count uint32                 `json:"count"`
	Cost  *ItemsCostServiceStats `json:"cost"`
}

type ItemsCostServiceStats struct {
	Avg int32 `json:"avg"`
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

type UsersServiceStats struct {
	Count   uint32                    `json:"count"`
	Credits *UsersCreditsServiceStats `json:"credits"`
}

type UsersCreditsServiceStats struct {
	Sum int32 `json:"sum"`
	Avg int32 `json:"avg"`
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

func newServiceStats(stats *matomat.ServiceStats) *ServiceStats {
	itemsStats := stats.Items
	itemsStatsCost := itemsStats.Cost
	usersStats := stats.Users
	usersStatsCredits := usersStats.Credits

	apiItemsCost := &ItemsCostServiceStats{Avg: itemsStatsCost.Avg, Min: itemsStatsCost.Min, Max: itemsStatsCost.Max}
	apiUsersCredits := &UsersCreditsServiceStats{Sum: usersStatsCredits.Sum, Avg: usersStatsCredits.Avg, Min: usersStatsCredits.Min, Max: usersStatsCredits.Max}

	apiItems := &ItemsServiceStats{Count: itemsStats.Count, Cost: apiItemsCost}
	apiUsers := &UsersServiceStats{Count: usersStats.Count, Credits: apiUsersCredits}

	apiStats := &ServiceStats{Items: apiItems, Users: apiUsers}

	return apiStats
}
