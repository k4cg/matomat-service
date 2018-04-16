package users

import (
	"errors"
)

type UserRepoMem struct {
	autoIncrement uint32
	users         map[uint32]User
	usersDeleted  map[uint32]User
}

const ERROR_UNKOWN_USER string = "Unkown user"

func NewUserRepoMem() *UserRepoMem {
	return &UserRepoMem{users: make(map[uint32]User), usersDeleted: make(map[uint32]User), autoIncrement: 1}
}

func (r *UserRepoMem) Get(id uint32) (User, error) {
	var err error
	user, found := r.users[id]
	if !found {
		err = errors.New(ERROR_UNKOWN_USER)
	}
	return user, err
}

func (r *UserRepoMem) GetByUsername(username string) (User, error) {
	var err error
	var user User
	for _, v := range r.users {
		if v.Username == username {
			user = v
			break
		}
	}
	if user == (User{}) {
		err = errors.New(ERROR_UNKOWN_USER)
	}
	return user, err
}

func (r *UserRepoMem) List() map[uint32]User {
	return r.users
}

func (r *UserRepoMem) Save(user User) (User, error) {
	var err error
	if user.ID != 0 {
		r.users[user.ID] = user
	} else {
		user.ID = r.autoIncrement
		r.autoIncrement = r.autoIncrement + 1
		r.users[user.ID] = user
	}
	return user, err
}

func (r *UserRepoMem) Delete(id uint32) (User, error) {
	user, found := r.users[id]
	if found {
		delete(r.users, id)
		r.usersDeleted[id] = user
	}
	return user, nil
}
