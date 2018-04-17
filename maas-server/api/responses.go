package api

import (
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/users"
)

type Error struct {
	Message string `json:"message"`
}

func newUser(user users.User) User {
	return User{ID: user.ID, Username: user.Username, Credits: user.Credits}
}

//Represenation of a Matomat users on API level
type User struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Credits  uint32 `json:"credits"`
}

type TransferredCredits struct {
	Sender  User   `json:"sender"`
	Credits uint32 `json:"credits"`
}

func newTransferredCredits(fromUser users.User, transferredCreditsAmount uint32) TransferredCredits {
	return TransferredCredits{Sender: newUser(fromUser), Credits: transferredCreditsAmount}
}

//Item Represenation of a drink/pizza/thing that is being "sold" by Matomat, containing its name and cost (in credits).
type Item struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Cost uint32 `json:"cost"`
}

func newItem(item items.Item) Item {
	return Item{ID: item.ID, Name: item.Name, Cost: item.Cost}
}

type ItemStats struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Cost     uint32 `json:"cost"`
	Consumed uint32 `json:"consumed"`
}

func newItemStats(item items.Item, itemStats items.ItemStats) ItemStats {
	return ItemStats{ID: item.ID, Name: item.Name, Cost: item.Cost, Consumed: itemStats.Consumed}
}

type AuthSuccess struct {
	Token               string `json:"token"`
	ExpirationTimestamp uint32 `json:"expires"`
}

func newAuthSuccess(token string, expirationTimestamp uint32) AuthSuccess {
	return AuthSuccess{Token: token, ExpirationTimestamp: expirationTimestamp}
}
