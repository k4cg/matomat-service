package api

import (
	"strings"

	"github.com/k4cg/matomat-service/maas-server/auth"
)

func BuildUsersRoutes(auth auth.AuthInterface, handler *UsersApiHandler) []Route {
	var routes = Routes{
		Route{
			"UsersGet",
			strings.ToUpper("Get"),
			"/v0/users",
			AuthenticationMiddleware(auth, handler.UsersGet),
		},

		Route{
			"UsersPost",
			strings.ToUpper("Post"),
			"/v0/users",
			AuthenticationMiddleware(auth, handler.UsersPost),
		},

		Route{
			"UsersUseridCreditsTransferPut",
			strings.ToUpper("Put"),
			"/v0/users/{userid}/credits/transfer",
			AuthenticationMiddleware(auth, handler.UsersUseridCreditsTransferPut),
		},

		Route{
			"UsersUseridDelete",
			strings.ToUpper("Delete"),
			"/v0/users/{userid}",
			AuthenticationMiddleware(auth, handler.UsersUseridDelete),
		},

		Route{
			"UsersUseridGet",
			strings.ToUpper("Get"),
			"/v0/users/{userid}",
			AuthenticationMiddleware(auth, handler.UsersUseridGet),
		},

		Route{
			"UsersUseridCreditsAddPut",
			strings.ToUpper("Put"),
			"/v0/users/{userid}/credits/add",
			AuthenticationMiddleware(auth, handler.UsersUseridCreditsAddPut),
		},

		Route{
			"UsersUseridCreditsWithdrawPut",
			strings.ToUpper("Put"),
			"/v0/users/{userid}/credits/withdraw",
			AuthenticationMiddleware(auth, handler.UsersUseridCreditsWithdrawPut),
		},

		Route{
			"UsersUseridStatsGet",
			strings.ToUpper("Get"),
			"/v0/users/{userid}/stats",
			AuthenticationMiddleware(auth, handler.UsersUseridStatsGet),
		},

		Route{
			"UsersPasswordPut",
			strings.ToUpper("Put"),
			"/v0/users/password",
			AuthenticationMiddleware(auth, handler.UsersPasswordPut),
		},

		Route{
			"UsersUserIdPasswordPut",
			strings.ToUpper("Put"),
			"/v0/users/{userid}/password",
			AuthenticationMiddleware(auth, handler.UsersUseridPasswordPut),
		},
	}
	return routes
}
