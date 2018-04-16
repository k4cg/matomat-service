/*
General TODO:
There sees to be too much boilderplate code. Solve this more cleverly (deepen understanding of golang for better constructs!)
*/
package api

import (
	"net/http"

	"github.com/k4cg/matomat-service/maas-server/users"

	"github.com/k4cg/matomat-service/maas-server/auth"
)

type AuthApiHandler struct {
	auth  auth.AuthInterface
	users *users.Users
}

const ERROR_UNAUTHORIZED string = "unauthorized"

func NewAuthApiHandler(auth auth.AuthInterface, users *users.Users) *AuthApiHandler {
	return &AuthApiHandler{auth: auth, users: users}
}

func extractLoginData(r *http.Request) (string, string, error) {
	var err error

	r.ParseForm()

	userName, err := formGet(r, FORM_KEY_USERNAME)
	userPassword, err := formGet(r, FORM_KEY_PASSWORD)

	return userName, userPassword, err
}

func (aah *AuthApiHandler) AuthLoginPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := http.StatusUnauthorized, []byte(ERROR_UNAUTHORIZED)

	userName, userPassword, err := extractLoginData(r)
	if err == nil {
		users, err := aah.users.IsPasswordValid(userName, userPassword)
		if err == nil {
			claims := aah.auth.NewAuthTokenClaims(users.ID, users.Username)
			token, err := aah.auth.GetToken(claims)
			if err == nil {
				status = http.StatusOK
				response = []byte(token) //TODO properly wrap the response / pack it in the right place (headers?)
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}
