package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	userRepo            UserRepositoryInterface
	passwordHashingCost int
}

const ERROR_CREATE_USER_USERNAME_ALREADY_TAKEN = "Username already taken"
const ERROR_PASSWORDS_DO_NOT_MATCH = "Passwords do not match"
const ERROR_CHANGE_PASSWORD_UNKNOWN_USERNAME = "Unkown username"
const ERROR_INVALID_USERNAME_OR_PASSWORD = "Invalid username or password"

func NewUsers(userRepo UserRepositoryInterface, passwordHashingCost int) *Users {
	return &Users{userRepo: userRepo, passwordHashingCost: passwordHashingCost}
}

func (ua *Users) GetUser(userId uint32) (User, error) {
	return ua.userRepo.Get(userId)
}

func (ua *Users) ListUsers() map[uint32]User {
	return ua.userRepo.List()
}

func (ua *Users) DeleteUser(userID uint32) (User, error) {
	return ua.userRepo.Delete(userID)
}

func (ua *Users) CreateUser(username string, password string, passwordRepeat string) (User, error) {
	var newUser User
	existingUser, err := ua.userRepo.GetByUsername(username)
	if err == nil {
		if existingUser == (User{}) {
			if password == passwordRepeat {
				hashedPassword, err := ua.hashPassword(password)
				if err == nil {
					newUser.Username = username
					newUser.Password = hashedPassword
					newUser, err = ua.userRepo.Save(newUser)
				}
			} else {
				err = errors.New(ERROR_PASSWORDS_DO_NOT_MATCH)
			}
		} else {
			err = errors.New(ERROR_CREATE_USER_USERNAME_ALREADY_TAKEN)
		}
	}
	return newUser, err
}

func (ua *Users) SetPassword(userId uint32, newPassword string, newPasswordRepeat string) (User, error) {
	user, err := ua.userRepo.Get(userId)
	if err == nil {
		if newPassword == newPasswordRepeat {
			hashedPassword, err := ua.hashPassword(newPassword)
			if err == nil {
				user.Password = hashedPassword
				user, err = ua.userRepo.Save(user)
			}
		} else {
			err = errors.New(ERROR_PASSWORDS_DO_NOT_MATCH)
		}
	}
	return user, err
}

func (ua *Users) ChangePassword(userId uint32, oldPassword string, newPassword string, newPasswordRepeat string) (User, error) {
	user, err := ua.userRepo.Get(userId)
	if err == nil {
		if ua.checkPasswordHash(oldPassword, user.Password) {
			if newPassword == newPasswordRepeat {
				hashedPassword, err := ua.hashPassword(newPassword)
				if err == nil {
					user.Password = hashedPassword
					user, err = ua.userRepo.Save(user)
				}
			} else {
				err = errors.New(ERROR_PASSWORDS_DO_NOT_MATCH)
			}
		} else {
			err = errors.New(ERROR_CHANGE_PASSWORD_UNKNOWN_USERNAME)
		}
	}
	return user, err
}

func (ua *Users) IsPasswordValid(username string, password string) (User, error) {
	var validUser User
	userAuth, err := ua.userRepo.GetByUsername(username)
	if err == nil {
		if ua.checkPasswordHash(password, userAuth.Password) {
			validUser = userAuth
		} else {
			err = errors.New(ERROR_INVALID_USERNAME_OR_PASSWORD)
		}
	}
	return validUser, err
}

func (ua *Users) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), ua.passwordHashingCost)
	return string(bytes), err
}

func (ua *Users) checkPasswordHash(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
