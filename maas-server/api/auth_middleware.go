package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/context"
	"github.com/k4cg/matomat-service/maas-server/auth"
	"strings"
)

const HEADER_AUTHORIZATION_KEY = "Authorization"
const ERROR_AUTHMIDDLEWARE_TOKEN_INVALID_AUTHORIZATION_TOKEN = "Invalid authorization token"
const ERROR_AUTHMIDDLEWARE_AUTHORIZATION_HEADER_REQUIRED = "An authorization header is required"
const ERROR_AUTHMIDDLEWARE_INVALID_AUTHORIZATION_DATA = "Invalid authorization data"
const CONTEXT_AUTHTOKENCLAIMS_USERID_KEY = "03u5rsx_dlFfh9sw"

func AuthenticationMiddleware(auth auth.AuthInterface, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerTokenHeaderString := req.Header.Get(HEADER_AUTHORIZATION_KEY)
		bearerToken, _ := extractBearerToken(bearerTokenHeaderString)

		if len(bearerToken) > 0 {
			authTokenClaims, err := auth.VerifyToken(bearerToken)
			if err != nil {
				w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
				w.WriteHeader(http.StatusBadRequest)
				response, _ := json.Marshal(Error{Message: err.Error()})
				w.Write(response)
				return
			} else {
				context.Set(req, CONTEXT_AUTHTOKENCLAIMS_USERID_KEY, authTokenClaims.ID)
				next.ServeHTTP(w, req)
			}
		} else {
			w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
			w.WriteHeader(http.StatusUnauthorized)
			response, _ := json.Marshal(Error{ERROR_AUTHMIDDLEWARE_AUTHORIZATION_HEADER_REQUIRED})
			w.Write(response)
			return
		}
	})
}

func extractBearerToken(bearerTokenHeaderString string) (string, error) {
	var bearerToken string
	var err error

	bearerTokenHeaderStringAr := strings.Fields(bearerTokenHeaderString)
	if len(bearerTokenHeaderStringAr) == 2 && strings.EqualFold(bearerTokenHeaderStringAr[0], "Bearer") {
		bearerToken = bearerTokenHeaderStringAr[1]
	} else {
		err = errors.New(ERROR_AUTHMIDDLEWARE_INVALID_AUTHORIZATION_DATA)
	}

	return bearerToken, err
}

func getUserIDFromContext(req *http.Request) (uint32, error) {
	var userID uint32
	var err error

	raw := context.Get(req, CONTEXT_AUTHTOKENCLAIMS_USERID_KEY)
	rawInt, ok := raw.(uint32)
	if ok {
		userID = rawInt
	} else {
		err = errors.New("Could not get user ID from context")
	}

	return userID, err
}
