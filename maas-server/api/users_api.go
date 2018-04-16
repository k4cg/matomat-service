/*
General TODO:
There sees to be too much boilderplate code. Solve this more cleverly (deepen understanding of golang for better constructs!)
*/
package api

import (
	"encoding/json"
	"net/http"

	"github.com/k4cg/matomat-service/maas-server/users"

	"github.com/k4cg/matomat-service/maas-server/auth"
	"github.com/k4cg/matomat-service/maas-server/matomat"
)

type UsersApiHandler struct {
	auth    auth.AuthInterface
	users   *users.Users
	matomat *matomat.Matomat
}

const FORM_KEY_CREDITS string = "credits"
const FORM_KEY_COST string = "cost"

func NewUsersApiHandler(auth auth.AuthInterface, users *users.Users, matomat *matomat.Matomat) *UsersApiHandler {
	return &UsersApiHandler{auth: auth, users: users, matomat: matomat}
}

func extractUserCreateData(r *http.Request) (string, string, string, error) {
	var err error

	r.ParseForm()

	userName, err := formGet(r, FORM_KEY_NAME)
	password, err := formGet(r, FORM_KEY_PASSWORD)
	passwordRepeat, err := formGet(r, FORM_KEY_PASSWORD_REPEAT)

	return userName, password, passwordRepeat, err
}

func extractPasswordChangeData(r *http.Request) (string, string, string, error) {
	var err error

	r.ParseForm()

	password, err := formGet(r, FORM_KEY_PASSWORD)
	newPassword, err := formGet(r, FORM_KEY_PASSWORD)
	newPasswordRepeat, err := formGet(r, FORM_KEY_PASSWORD_REPEAT)

	return password, newPassword, newPasswordRepeat, err
}

func extractPasswordChangeUseridData(r *http.Request) (string, string, error) {
	var err error

	r.ParseForm()

	newPassword, err := formGet(r, FORM_KEY_PASSWORD)
	newPasswordRepeat, err := formGet(r, FORM_KEY_PASSWORD_REPEAT)

	return newPassword, newPasswordRepeat, err
}

func extractUserCreditsChangeData(r *http.Request) (int32, error) {
	var err error

	r.ParseForm()

	userCredits, err := formGetInt32(r, FORM_KEY_CREDITS)

	return userCredits, err
}

func (uah *UsersApiHandler) UsersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_USERID_GET) {
		users := uah.users.ListUsers()
		apiUsers := make(map[uint32]User)
		for k, v := range users {
			apiUsers[k] = newUser(v)
		}
		status, response = getResponse(http.StatusOK, apiUsers)
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_USERID_GET) {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil {
			user, err := uah.users.GetUser(userID)
			if err == nil {
				status, response = getResponse(http.StatusOK, newUser(user))
			} else {
				status, response = getErrorResponse(http.StatusNotFound, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_CREATE) {
		name, password, passwordRepeat, err := extractUserCreateData(r)
		if err == nil {
			user, err := uah.users.CreateUser(name, password, passwordRepeat)
			if err == nil {
				status, response = getResponse(http.StatusCreated, newUser(user))
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_USERID_DELETE) {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil {
			user, err := uah.users.DeleteUser(userID)
			if err == nil {
				status, response = getResponse(http.StatusOK, newUser(user))
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridCreditsAddPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	userID, err := extractIDFromModelGet(r.URL.Path)
	loggedInUserID := getUserIDFromContext(r)

	if uah.matomat.IsAllowed(loggedInUserID, matomat.ACTION_USERS_USERID_CREDITS_ADD) && err == nil && loggedInUserID == userID {
		credits, err := extractUserCreditsChangeData(r)
		if err == nil {
			user, err := uah.matomat.UserCreditsAdd(userID, credits)
			if err == nil {
				status, response = getResponse(http.StatusOK, newUser(user))
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridCreditsWithdrawPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	userID, err := extractIDFromModelGet(r.URL.Path)
	loggedInUserID := getUserIDFromContext(r)

	if uah.matomat.IsAllowed(loggedInUserID, matomat.ACTION_USERS_USERID_CREDITS_WITHDRAW) && err == nil && loggedInUserID == userID {
		credits, err := extractUserCreditsChangeData(r)
		if err == nil {
			user, err := uah.matomat.UserCreditsWithdraw(userID, credits)
			if err == nil {
				status, response = getResponse(http.StatusOK, newUser(user))
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridStatsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_USERID_STATS_GET) {
		status = http.StatusNotImplemented
		response, _ = json.Marshal(ERROR_NOT_IMPLEMENTED) //TODO implement!
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridCreditsTransferPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()

	userID, err := extractIDFromModelGet(r.URL.Path)
	loggedInUserID := getUserIDFromContext(r)

	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_USERID_CREDITS_TRANSFER) && err == nil && userID != 0 {
		credits, err := extractUserCreditsChangeData(r)
		if err == nil {
			fromUser, transferredCredits, err := uah.matomat.CreditsTransfer(loggedInUserID, userID, credits)
			if err == nil {
				status, response = getResponse(http.StatusOK, newTransferredCredits(fromUser, transferredCredits))
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersPasswordPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := http.StatusUnauthorized, []byte(ERROR_UNAUTHORIZED)

	oldPassword, newPassword, newPasswordRepeat, err := extractPasswordChangeData(r)
	if err == nil && uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_OWN_PASSWORD_CHANGE) {
		uah.users.ChangePassword(getUserIDFromContext(r), oldPassword, newPassword, newPasswordRepeat)
	} else {
		status, response = getErrorResponse(http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridPasswordPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := http.StatusUnauthorized, []byte(ERROR_UNAUTHORIZED)

	//explicitly NOT allow to set own users password using this method, to not have a way to sneak around changing the
	//"own" password without known the old one
	if uah.matomat.IsAllowed(getUserIDFromContext(r), matomat.ACTION_USERS_USERID_PASSWORD_CHANGE) {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil {
			newPassword, newPasswordRepeat, err := extractPasswordChangeUseridData(r)
			if err == nil {
				uah.users.SetPassword(userID, newPassword, newPasswordRepeat)
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}
