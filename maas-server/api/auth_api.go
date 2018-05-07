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

func extractLoginData(r *http.Request) (string, string, uint32, error) {
	var err error

	r.ParseForm()

	userName, err := formGet(r, FORM_KEY_USERNAME)
	userPassword, err := formGet(r, FORM_KEY_PASSWORD)
	tokenValiditySeconds, _ := formGetInt32(r, FORM_KEY_TOKEN_VALIDITY_SECONDS)

	//TODO again, another evil cast ...
	return userName, userPassword, uint32(tokenValiditySeconds), err
}

func (aah *AuthApiHandler) AuthLoginPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := http.StatusUnauthorized, []byte(ERROR_UNAUTHORIZED)

	userName, userPassword, tokenValiditySeconds, err := extractLoginData(r)
	if err == nil {
		user, err := aah.users.IsPasswordValid(userName, userPassword)
		if err == nil {
			claims := aah.auth.NewAuthTokenClaims(user.ID, user.Username, tokenValiditySeconds)
			token, err := aah.auth.GetToken(claims)
			if err == nil {
				status, response = getResponse(http.StatusOK, newAuthSuccess(token, uint32(claims.ExpiresAt), user)) //TODO another evil cast...
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
		}
	} else {
		status, response = getErrorResponse(http.StatusUnauthorized, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}
