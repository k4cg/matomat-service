package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/k4cg/matomat-service/maas-server/auth"
)

const HEADER_AUTHORIZATION_KEY = "Authorization"
const ERROR_AUTHMIDDLEWARE_TOKEN_INVALID_AUTHORIZATION_TOKEN = "Invalid authorization token"
const ERROR_AUTHMIDDLEWARE_AUTHORIZATION_HEADER_REQUIRED = "An authorization header is required"
const ERROR_AUTHMIDDLEWARE_INVALID_AUTHORIZATION_DATA = "Invalid authorization data"
const CONTEXT_AUTHTOKENCLAIMS_USERID_KEY = "03u5rsx_dlFfh9sw"

func AuthenticationMiddleware(auth auth.AuthInterface, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerToken := req.Header.Get(HEADER_AUTHORIZATION_KEY)
		if bearerToken != "" {
			authTokenClaims, err := auth.VerifyToken(bearerToken)
			if err != nil {
				w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
				w.WriteHeader(http.StatusBadRequest)
				response, _ := json.Marshal(Error{Message: err.Error()})
				w.Write(response)
				return
			} else {
				userIDString := strconv.Itoa(int(authTokenClaims.ID))
				//TODO THIS IS ALL WONKY! reimplement properly ?? perhaps store the object in there?
				context.Set(req, CONTEXT_AUTHTOKENCLAIMS_USERID_KEY, []byte(userIDString))
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

func getUserIDFromContext(req *http.Request) uint32 {
	var userID uint32
	//TODO THIS IS ALL WONKY! reimplement properly ?? perhaps store the object in there?
	raw := context.Get(req, CONTEXT_AUTHTOKENCLAIMS_USERID_KEY)
	userIDString, ok := raw.(string)
	if ok {
		userIDInt, err := strconv.Atoi(userIDString)
		if err != nil {
			userID = uint32(userIDInt)
		}
	}

	return userID
}
