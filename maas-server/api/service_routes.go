package api

import (
	"strings"

	"github.com/k4cg/matomat-service/maas-server/auth"
)

func BuildServiceRoutes(auth auth.AuthInterface, handler *ServiceApiHandler) []Route {
	var routes = Routes{
		Route{
			"ServiceStatsGet",
			strings.ToUpper("Get"),
			"/v0/service/stats",
			AuthenticationMiddleware(auth, handler.ServiceStatsGet),
		},
	}
	return routes
}
