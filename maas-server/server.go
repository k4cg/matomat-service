/*
 * MaaS - Server
 */

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olebedev/config"

	"github.com/k4cg/matomat-service/maas-server/api"
	"github.com/k4cg/matomat-service/maas-server/auth"
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/matomat"
	"github.com/k4cg/matomat-service/maas-server/users"
)

const CONFIG_FILE_PATH = "config.yml"

func buildRepos(cfg *config.Config) (users.UserRepositoryInterface, items.ItemRepositoryInterface, items.ItemStatsRepositoryInterface) {
	//TODO add error handling / checking on config value retrieval
	userRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.users")
	itemRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.items")
	itemStatsRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.items_stats")

	userRepo := users.NewUserRepoSqlite3(userRepoSqlite3DbFilePath)
	itemRepo := items.NewItemRepoSqlite3(itemRepoSqlite3DbFilePath)
	itemStatsRepo := items.NewItemStatsRepoSqlite3(itemStatsRepoSqlite3DbFilePath)

	return userRepo, itemRepo, itemStatsRepo
}

func buildApiHandlers(auth auth.AuthInterface, users *users.Users, matomat *matomat.Matomat) (*api.AuthApiHandler, *api.UsersApiHandler, *api.ItemsApiHandler) {
	authApiHandler := api.NewAuthApiHandler(auth, users)
	usersApiHandler := api.NewUsersApiHandler(auth, users, matomat)
	itemsApiHandler := api.NewItemsApiHandler(auth, matomat)

	return authApiHandler, usersApiHandler, itemsApiHandler
}

func buildRoutes(auth auth.AuthInterface, authApiHandler *api.AuthApiHandler, usersApiHandler *api.UsersApiHandler, itemsApiHandler *api.ItemsApiHandler) []api.Route {
	authRoutes := api.BuildAuthRoutes(auth, authApiHandler)
	usersRoutes := api.BuildUsersRoutes(auth, usersApiHandler)
	itemsRoutes := api.BuildItemsRoutes(auth, itemsApiHandler)

	routes := append(itemsRoutes, usersRoutes...)
	return append(routes, authRoutes...)
}

func buildAuth(cfg *config.Config) *auth.AuthJWT {
	//TODO add error handling / checking on config value retrieval
	issuer, _ := cfg.String("jwt.issuer")
	secret, _ := cfg.String("jwt.sig.secret")
	secondsValid, _ := cfg.Int("jwt.valid_sec")
	return auth.NewAuthJWT(issuer, secret, uint32(secondsValid))
}

func buildUsers(cfg *config.Config, userRepo users.UserRepositoryInterface) *users.Users {
	//TODO add error handling / checking on config value retrieval
	hashRounds, _ := cfg.Int("auth.hash.rounds")
	return users.NewUsers(userRepo, hashRounds)
}

func runServer(cfg *config.Config, router *mux.Router) error {
	//TODO add error handling / checking on config value retrieval
	addr, _ := cfg.String("listen.addr")
	port, _ := cfg.String("listen.port")
	sslServerKeyFilePath, _ := cfg.String("ssl.key")
	sslServerCertFilePath, _ := cfg.String("ssl.cert")

	log.Printf("MaaS server started at " + addr + ":" + port)
	return http.ListenAndServeTLS(addr+":"+port, sslServerCertFilePath, sslServerKeyFilePath, router)
}

func main() {
	cfg, err := config.ParseYamlFile(CONFIG_FILE_PATH)
	if err == nil {
		userRepo, itemRepo, itemStatsRepo := buildRepos(cfg)
		auth := buildAuth(cfg)
		users := buildUsers(cfg, userRepo)

		eventDispatcherMqtt := matomat.NewEventDispatcherMqtt() //TODO intialize properly when implemented
		matomat := matomat.NewMatomat(eventDispatcherMqtt, userRepo, itemRepo, itemStatsRepo)

		authApiHandler, usersApiHandler, itemsApiHandler := buildApiHandlers(auth, users, matomat)

		routes := buildRoutes(auth, authApiHandler, usersApiHandler, itemsApiHandler)

		router := api.NewRouter(routes)

		log.Fatal(runServer(cfg, router))
	} else {
		log.Fatal(err)
	}
}
