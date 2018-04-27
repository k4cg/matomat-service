package matomat

import (
	"errors"

	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/users"
)

type Matomat struct {
	eventDispatcher    EventDispatcherInterface
	userRepo           users.UserRepositoryInterface
	itemRepo           items.ItemRepositoryInterface
	itemStatsRepo      items.ItemStatsRepositoryInterface
	userItemsStatsRepo users.UserItemsStatsRepositoryInterface
}

//TODO is this the right place to put the actions? should they be moved out?
const ACTION_ITEMS_ITEMID_CONSUME string = "ItemsItemidConsumePut"
const ACTION_ITEMS_ITEMID_DELETE string = "ItemsItemidDelete"
const ACTION_ITEMS_ITEMID_GET string = "ItemsItemidGet"
const ACTION_ITEMS_ITEMID_EDIT string = "ItemsItemidPut"
const ACTION_ITEMS_ITEMID_STATS_GET string = "ItemsItemidStatsGet"
const ACTION_ITEMS_LIST string = "ItemsGet"
const ACTION_ITEMS_CREATE string = "ItemsPost"
const ACTION_ITEMS_STATS_GET string = "ItemsStatsGet"

const ACTION_USERS_LIST string = "UsersGet"
const ACTION_USERS_CREATE string = "UsersPost"
const ACTION_USERS_USERID_CREDITS_TRANSFER string = "UsersUseridCreditsTransferPut"
const ACTION_USERS_USERID_DELETE string = "UsersUseridDelete"
const ACTION_USERS_USERID_GET string = "UsersUseridGet"
const ACTION_USERS_USERID_CREDITS_ADD string = "UsersUseridCreditsAddPut"
const ACTION_USERS_USERID_CREDITS_WITHDRAW string = "UsersUseridCreditsWithdrawPut"
const ACTION_USERS_USERID_STATS_GET string = "UsersUseridStatsGet"
const ACTION_USERS_OWN_PASSWORD_CHANGE string = "UserPasswordPut"
const ACTION_USERS_USERID_PASSWORD_CHANGE string = "UsersUseridPasswordPut"

const ERROR_CONSUME_ITEM_NOT_ENOUGH_CREDITS string = "Not enough credits"
const ERROR_TRANSFER_CREDITS_NOT_ENOUGH_CREDITS string = "User does not have enough credits to transfer"
const ERROR_TRANSFER_UNKOWN_CREDITS_RECEIVER string = "User to transfer credits to not unknown"
const ERROR_TRANSFER_UNKOWN_CREDITS_SENDER string = "User to transfer credits from unknown"
const ERROR_USER_CREDITS_WITHDRAW_NOT_ENOUGH_CREDITS string = "Not enough credits. Cannot withdraw more credits than the current balance"
const ERROR_USER_ONLY_POSITIVE_OR_ZERO_CREDIT_VALUES_ALLOWED string = "Only credit amounts greater or equal to zero allowed"
const ERROR_UNKNOWN_ITEM string = "Unkown item"
const ERROR_UNKNOWN_USER string = "Unkown user"
const ERROR_UNKNOWN_USER_FROM string = "Unkown receiving user"
const ERROR_UNKNOWN_USER_TO string = "Unkown receiving user"

func NewMatomat(eventDispatcher EventDispatcherInterface, userRepo users.UserRepositoryInterface, itemRepo items.ItemRepositoryInterface, itemStatsRepo items.ItemStatsRepositoryInterface, userItemsStatsRepo users.UserItemsStatsRepositoryInterface) *Matomat {
	return &Matomat{eventDispatcher: eventDispatcher, userRepo: userRepo, itemRepo: itemRepo, itemStatsRepo: itemStatsRepo, userItemsStatsRepo: userItemsStatsRepo}
}

func (m *Matomat) IsAllowed(userID uint32, action string) bool {
	allowed := false

	adminRequiredActions := make(map[string]string)
	adminRequiredActions[ACTION_USERS_CREATE] = ACTION_USERS_CREATE
	adminRequiredActions[ACTION_USERS_USERID_DELETE] = ACTION_USERS_USERID_DELETE
	adminRequiredActions[ACTION_USERS_USERID_PASSWORD_CHANGE] = ACTION_USERS_USERID_PASSWORD_CHANGE

	_, requiresAdmin := adminRequiredActions[action]

	if requiresAdmin {
		user, err := m.userRepo.Get(userID)
		if err == nil {
			if user != (users.User{}) {
				allowed = user.IsAdmin()
			}
		} //if user is not found (or any error occurs) => access is forbidden
	} else {
		allowed = true
	}

	return allowed
}

func (m *Matomat) ItemConsume(userID uint32, itemID uint32) (items.Item, items.ItemStats, error) {
	var remainingCredits uint32
	var itemStatsToReturn items.ItemStats
	var retErr error

	item, err := m.itemRepo.Get(itemID)
	if err == nil {
		if item != (items.Item{}) {
			user, err := m.userRepo.Get(userID)
			if err == nil {
				if user != (users.User{}) {
					if user.Credits >= item.Cost {
						remainingCredits = user.Credits - item.Cost
						user.Credits = remainingCredits
						m.userRepo.Save(user)
						go m.itemStatsRepo.CountConsumption(item.ID, 1)
						go m.userItemsStatsRepo.CountConsumption(userID, item.ID, 1)
						go m.eventDispatcher.ItemConsumed(user.ID, user.Username, item.ID, item.Name, item.Cost)
						itemStats, err := m.itemStatsRepo.Get(itemID)
						if err == nil {
							itemStatsToReturn = itemStats
						} else {
							retErr = err
						}
					} else {
						retErr = errors.New(ERROR_CONSUME_ITEM_NOT_ENOUGH_CREDITS)
					}
				} else {
					retErr = errors.New(ERROR_UNKNOWN_USER)
				}
			} else {
				retErr = err
			}
		} else {
			retErr = errors.New(ERROR_UNKNOWN_ITEM)
		}
	} else {
		retErr = err
	}

	return item, itemStatsToReturn, retErr
}

func (m *Matomat) CreditsTransfer(fromUserID uint32, toUserID uint32, amountToTransfer int32) (users.User, uint32, error) {
	var transferredAmount uint32
	var oldFromCredits uint32
	var retErr error
	var senderToReturn users.User

	if amountToTransfer >= 0 {
		amount := uint32(amountToTransfer) //TODO those "blind" uint32 casts should probably be handled better...
		fromUser, err := m.userRepo.Get(fromUserID)
		if err == nil {
			if fromUser != (users.User{}) {
				toUser, err := m.userRepo.Get(toUserID)
				if err == nil {
					if toUser != (users.User{}) {
						if fromUser.Credits >= amount {
							//yes this is not "transaction save" ... feel free to improve :-P ... e.g. move to a separate repo call
							fromUser.Credits = fromUser.Credits - amount
							toUser.Credits = toUser.Credits + amount
							fromUser, err = m.userRepo.Save(fromUser)
							if err == nil {
								toUser, err = m.userRepo.Save(toUser)
								if err != nil {
									//"ROLLBACK"
									fromUser.Credits = oldFromCredits
									fromUser, err = m.userRepo.Save(fromUser) //if this does not work ... we're fucked ^^
									retErr = err
								} else {
									transferredAmount = amount
								}
							} else {
								retErr = err
							}
							senderToReturn = fromUser
						} else {
							retErr = errors.New(ERROR_TRANSFER_CREDITS_NOT_ENOUGH_CREDITS)
						}
					} else {
						retErr = errors.New(ERROR_UNKNOWN_USER_TO)
					}
				} else {
					retErr = errors.New(ERROR_TRANSFER_UNKOWN_CREDITS_RECEIVER)
				}
			} else {
				retErr = errors.New(ERROR_UNKNOWN_USER_FROM)
			}
		} else {
			retErr = errors.New(ERROR_TRANSFER_UNKOWN_CREDITS_SENDER)
		}
	} else {
		retErr = errors.New(ERROR_USER_ONLY_POSITIVE_OR_ZERO_CREDIT_VALUES_ALLOWED)
	}
	return senderToReturn, transferredAmount, retErr
}

func (m *Matomat) ItemGet(itemID uint32) (items.Item, error) {
	return m.itemRepo.Get(itemID)
}

func (m *Matomat) ItemDelete(itemID uint32) (items.Item, error) {
	return m.itemRepo.Delete(itemID)
}

func (m *Matomat) ItemCreate(name string, cost int32) (items.Item, error) {
	var retItem items.Item
	var retErr error
	if cost >= 0 {
		item := items.Item{Name: name, Cost: uint32(cost)} //TODO those "blind" uint32 casts should probably be handled better...
		item, err := m.itemRepo.Save(item)
		retItem = item
		retErr = err
	} else {
		retErr = errors.New(ERROR_USER_ONLY_POSITIVE_OR_ZERO_CREDIT_VALUES_ALLOWED)
	}
	return retItem, retErr
}

func (m *Matomat) ItemUpdate(itemID uint32, name string, cost int32) (items.Item, error) {
	var returnedItem items.Item
	var retErr error
	if cost >= 0 {
		item, err := m.itemRepo.Get(itemID)
		if err == nil && item != (items.Item{}) {
			item.Name = name
			item.Cost = uint32(cost) //TODO those "blind" uint32 casts should probably be handled better...
			item, err = m.itemRepo.Save(item)
			returnedItem = item
			retErr = err
		} else {
			retErr = errors.New(ERROR_UNKNOWN_ITEM)
		}
	} else {
		retErr = errors.New(ERROR_USER_ONLY_POSITIVE_OR_ZERO_CREDIT_VALUES_ALLOWED)
	}
	return returnedItem, retErr
}

func (m *Matomat) ItemsList() ([]items.Item, error) {
	return m.itemRepo.List()
}

func (m *Matomat) ItemGetStats(itemID uint32) (items.Item, items.ItemStats, error) {
	var item items.Item
	var itemStats items.ItemStats
	var retErr error

	item, err := m.itemRepo.Get(itemID)
	if err == nil {
		if item != (items.Item{}) {
			itemStats, err = m.itemStatsRepo.Get(itemID)
			retErr = err
		} else {
			retErr = errors.New(ERROR_UNKNOWN_ITEM)
		}
	} else {
		retErr = err
	}

	return item, itemStats, retErr
}

func (m *Matomat) ItemStatsList() ([]items.ItemStats, error) {
	return m.itemStatsRepo.List()
}

func (m *Matomat) UserCreditsAdd(userID uint32, credits int32) (users.User, error) {
	var retErr error
	user, err := m.userRepo.Get(userID)

	if err == nil {
		if user != (users.User{}) {
			if credits >= 0 {
				user.Credits = user.Credits + uint32(credits) //TODO those "blind" uint32 casts should probably be handled better...
				user, err = m.userRepo.Save(user)
				retErr = err
			} else {
				retErr = errors.New(ERROR_USER_ONLY_POSITIVE_OR_ZERO_CREDIT_VALUES_ALLOWED)
			}
		} else {
			retErr = errors.New(ERROR_UNKNOWN_USER)
		}
	}

	return user, retErr
}

func (m *Matomat) UserCreditsWithdraw(userID uint32, credits int32) (users.User, error) {
	var retErr error
	user, err := m.userRepo.Get(userID)

	if err == nil {
		if user != (users.User{}) {
			if credits >= 0 {
				withdrawAmount := uint32(credits) //TODO those "blind" uint32 casts should probably be handled better...
				if user.Credits >= withdrawAmount {
					user.Credits = user.Credits - uint32(credits)
					user, err = m.userRepo.Save(user)
					retErr = err
				} else {
					retErr = errors.New(ERROR_USER_CREDITS_WITHDRAW_NOT_ENOUGH_CREDITS)
				}
			} else {
				retErr = errors.New(ERROR_USER_ONLY_POSITIVE_OR_ZERO_CREDIT_VALUES_ALLOWED)
			}
		} else {
			retErr = errors.New(ERROR_UNKNOWN_USER)
		}
	} else {
		retErr = err
	}

	return user, retErr
}

func (m *Matomat) UsersUseridStatsGet(userID uint32) ([]items.ItemStats, error) {
	return m.userItemsStatsRepo.Get(userID)
}
