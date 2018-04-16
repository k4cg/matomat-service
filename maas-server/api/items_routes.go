package api

import (
	"strings"

	"github.com/k4cg/matomat-service/maas-server/auth"
)

func BuildItemsRoutes(auth auth.AuthInterface, handler *ItemsApiHandler) []Route {
	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/v0/",
			Index,
		},

		Route{
			"ItemsItemidConsumePut",
			strings.ToUpper("Put"),
			"/v0/items/{itemid}/consume",
			AuthenticationMiddleware(auth, handler.ItemsItemidConsumePut),
		},

		Route{
			"ItemsItemidDelete",
			strings.ToUpper("Delete"),
			"/v0/items/{itemid}",
			AuthenticationMiddleware(auth, handler.ItemsItemidDelete),
		},

		Route{
			"ItemsItemidGet",
			strings.ToUpper("Get"),
			"/v0/items/{itemid}",
			AuthenticationMiddleware(auth, handler.ItemsItemidGet),
		},

		Route{
			"ItemsItemidPut",
			strings.ToUpper("Put"),
			"/v0/items/{itemid}",
			AuthenticationMiddleware(auth, handler.ItemsItemidPut),
		},

		Route{
			"ItemsItemidStatsGet",
			strings.ToUpper("Get"),
			"/v0/items/{itemid}/stats",
			AuthenticationMiddleware(auth, handler.ItemsItemidStatsGet),
		},

		Route{
			"ItemsGet",
			strings.ToUpper("Get"),
			"/v0/items",
			AuthenticationMiddleware(auth, handler.ItemsGet),
		},

		Route{
			"ItemsPost",
			strings.ToUpper("Post"),
			"/v0/items",
			AuthenticationMiddleware(auth, handler.ItemsPost),
		},

		Route{
			"ItemsStatsGet",
			strings.ToUpper("Get"),
			"/v0/items/stats",
			AuthenticationMiddleware(auth, handler.ItemsStatsGet),
		},
	}
	return routes
}
