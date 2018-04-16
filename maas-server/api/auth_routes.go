package api

import (
	"strings"

	"github.com/k4cg/matomat-service/maas-server/auth"
)

func BuildAuthRoutes(auth auth.AuthInterface, handler *AuthApiHandler) []Route {
	var routes = Routes{
		Route{
			"AuthLoginPost",
			strings.ToUpper("Post"),
			"/v0/auth/login",
			handler.AuthLoginPost,
		},
	}
	return routes
}
