/*
General TODO:
There sees to be too much boilderplate code. Solve this more cleverly (deepen understanding of golang for better constructs!)
*/
package api

import (
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
	newPassword, err := formGet(r, FORM_KEY_PASSWORD_NEW)
	newPasswordRepeat, err := formGet(r, FORM_KEY_PASSWORD_REPEAT)

	return password, newPassword, newPasswordRepeat, err
}

func extractPasswordChangeUseridData(r *http.Request) (string, string, error) {
	var err error

	r.ParseForm()

	newPassword, err := formGet(r, FORM_KEY_PASSWORD_NEW)
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
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_GET) {
			users, err := uah.users.ListUsers()
			if err == nil {
				apiUsers := make(map[uint32]User)
				for k, v := range users {
					apiUsers[k] = newUser(v)
				}
				status, response = getResponse(http.StatusOK, apiUsers)
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_GET) {
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
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_CREATE) {
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
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_DELETE) {
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
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridCreditsAddPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil && uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_CREDITS_ADD) && loginUserID == userID {
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
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridCreditsWithdrawPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil && uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_CREDITS_WITHDRAW) && loginUserID == userID {
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
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridStatsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil && uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_STATS_GET) && userID != 0 {
			items, err := uah.matomat.ItemsList()
			if err == nil {
				itemStatsList, err := uah.matomat.UsersUseridStatsGet(userID)
				if err == nil {
					apiItemStats := make(map[uint32]ItemStats)
					for k, v := range items {
						itemStats, found := itemStatsList[k]
						if found {
							apiItemStats[k] = newItemStats(v, itemStats)
						}
					}
					status, response = getResponse(http.StatusOK, apiItemStats)
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridCreditsTransferPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		userID, err := extractIDFromModelGet(r.URL.Path)
		if err == nil && uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_CREDITS_TRANSFER) && userID != 0 {
			credits, err := extractUserCreditsChangeData(r)
			if err == nil {
				fromUser, transferredCredits, err := uah.matomat.CreditsTransfer(loginUserID, userID, credits)
				if err == nil {
					status, response = getResponse(http.StatusOK, newTransferredCredits(fromUser, transferredCredits))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersPasswordPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := http.StatusUnauthorized, []byte(ERROR_UNAUTHORIZED)
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		oldPassword, newPassword, newPasswordRepeat, err := extractPasswordChangeData(r)
		if err == nil && uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_OWN_PASSWORD_CHANGE) {
			user, err := uah.users.ChangePassword(loginUserID, oldPassword, newPassword, newPasswordRepeat)
			if err == nil {
				status, response = getResponse(http.StatusOK, newUser(user))
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusBadRequest, err.Error())
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (uah *UsersApiHandler) UsersUseridPasswordPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := http.StatusUnauthorized, []byte(ERROR_UNAUTHORIZED)
	loginUserID, err := getUserIDFromContext(r)

	//explicitly NOT allow to set own users password using this method, to not have a way to sneak around changing the
	//"own" password without known the old one
	if err == nil {
		if uah.matomat.IsAllowed(loginUserID, matomat.ACTION_USERS_USERID_PASSWORD_CHANGE) {
			userID, err := extractIDFromModelGet(r.URL.Path)
			if err == nil {
				newPassword, newPasswordRepeat, err := extractPasswordChangeUseridData(r)
				if err == nil {
					user, err := uah.users.SetPassword(userID, newPassword, newPasswordRepeat)
					if err == nil {
						status, response = getResponse(http.StatusOK, newUser(user))
					} else {
						status, response = getErrorResponse(http.StatusBadRequest, err.Error())
					}
				} else {
					status, response = getErrorResponse(http.StatusBadRequest, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}
