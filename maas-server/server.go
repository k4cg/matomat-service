/*
 * MaaS - Server
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/olebedev/config"

	"github.com/k4cg/matomat-service/maas-server/api"
	"github.com/k4cg/matomat-service/maas-server/auth"
	"github.com/k4cg/matomat-service/maas-server/items"
	"github.com/k4cg/matomat-service/maas-server/matomat"
	"github.com/k4cg/matomat-service/maas-server/users"
)

const CONFIG_FILE_PATH = "config.yml"

func main() {
	cfg, err := config.ParseYamlFile(CONFIG_FILE_PATH)

	if err == nil {
		setupSignalHandling(cfg)

		userRepo, itemRepo, itemStatsRepo, userItemStatsRepo := buildRepos(cfg)
		auth := buildAuth(cfg)
		users := buildUsers(cfg, userRepo)

		eventDispatcherMqtt := buildEventDispatcher(cfg)
		matomatConfig := buildMatomatConfig(cfg)
		matomat := matomat.NewMatomat(*matomatConfig, eventDispatcherMqtt, userRepo, itemRepo, itemStatsRepo, userItemStatsRepo)

		authApiHandler, usersApiHandler, itemsApiHandler, serviceApiHandler := buildApiHandlers(auth, users, matomat)

		routes := buildRoutes(auth, authApiHandler, usersApiHandler, itemsApiHandler, serviceApiHandler)

		router := api.NewRouter(routes)

		log.Fatal(runServer(cfg, router))
	} else {
		log.Fatal(err)
	}
}

func buildRepos(cfg *config.Config) (users.UserRepositoryInterface, items.ItemRepositoryInterface, items.ItemStatsRepositoryInterface, users.UserItemsStatsRepositoryInterface) {
	//TODO add error handling / checking on config value retrieval
	userRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.users")
	itemRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.items")
	itemStatsRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.items_stats")
	userItemsStatsRepoSqlite3DbFilePath, _ := cfg.String("db.sqlite3.user_items_stats")

	userRepo := users.NewUserRepoSqlite3(userRepoSqlite3DbFilePath)
	itemRepo := items.NewItemRepoSqlite3(itemRepoSqlite3DbFilePath)
	itemStatsRepo := items.NewItemStatsRepoSqlite3(itemStatsRepoSqlite3DbFilePath)
	userItemStatsRepo := users.NewUserItemsStatsRepoSqlite3(userItemsStatsRepoSqlite3DbFilePath)

	return userRepo, itemRepo, itemStatsRepo, userItemStatsRepo
}

func buildApiHandlers(auth auth.AuthInterface, users *users.Users, matomat *matomat.Matomat) (*api.AuthApiHandler, *api.UsersApiHandler, *api.ItemsApiHandler, *api.ServiceApiHandler) {
	authApiHandler := api.NewAuthApiHandler(auth, users)
	usersApiHandler := api.NewUsersApiHandler(auth, users, matomat)
	itemsApiHandler := api.NewItemsApiHandler(auth, matomat)
	serviceApiHandler := api.NewServiceApiHandler(auth, matomat)
	return authApiHandler, usersApiHandler, itemsApiHandler, serviceApiHandler
}

func buildRoutes(auth auth.AuthInterface, authApiHandler *api.AuthApiHandler, usersApiHandler *api.UsersApiHandler, itemsApiHandler *api.ItemsApiHandler, serviceApiHandler *api.ServiceApiHandler) []api.Route {
	authRoutes := api.BuildAuthRoutes(auth, authApiHandler)
	usersRoutes := api.BuildUsersRoutes(auth, usersApiHandler)
	itemsRoutes := api.BuildItemsRoutes(auth, itemsApiHandler)
	serviceRoutes := api.BuildServiceRoutes(auth, serviceApiHandler)
	routes := append(itemsRoutes, usersRoutes...)
	routes = append(routes, authRoutes...)
	return append(routes, serviceRoutes...)
}

func buildAuth(cfg *config.Config) *auth.AuthJWT {
	//TODO add error handling / checking on config value retrieval
	issuer, _ := cfg.String("jwt.issuer")
	secret, _ := cfg.String("jwt.sig.secret")
	secondsValidDefault, _ := cfg.Int("jwt.valid_sec.default")
	secondsValidMax, _ := cfg.Int("jwt.valid_sec.max")
	return auth.NewAuthJWT(issuer, secret, uint32(secondsValidDefault), uint32(secondsValidMax))
}

func buildUsers(cfg *config.Config, userRepo users.UserRepositoryInterface) *users.Users {
	//TODO add error handling / checking on config value retrieval
	hashRounds, _ := cfg.Int("auth.hash.rounds")
	return users.NewUsers(userRepo, hashRounds)
}

func buildEventDispatcher(cfg *config.Config) matomat.EventDispatcherInterface {
	//TODO add error handling / checking on config value retrieval
	clientID, _ := cfg.String("event_dispatching.mqtt.client_id")
	connectionString, _ := cfg.String("event_dispatching.mqtt.connection_string")
	topic, _ := cfg.String("event_dispatching.mqtt.topic")
	enabled, _ := cfg.Bool("event_dispatching.enabled")
	return matomat.NewEventDispatcherMqtt(connectionString, clientID, topic, enabled)
}

func buildMatomatConfig(cfg *config.Config) *matomat.Config {
	//TODO add error handling / checking on config value retrieval
	allowDebt, _ := cfg.Bool("application.credit.allow_debt")
	itemNameMinLength, _ := cfg.Int("application.item.name_min_length")
	itemNameMaxLength, _ := cfg.Int("application.item.name_max_length")
	return matomat.NewConfig(allowDebt, uint32(itemNameMinLength), uint32(itemNameMaxLength))
}

func runServer(cfg *config.Config, router *mux.Router) error {
	//TODO add error handling / checking on config value retrieval
	addr, _ := cfg.String("listen.addr")
	port, _ := cfg.String("listen.port")
	sslServerKeyFilePath, _ := cfg.String("ssl.key")
	sslServerCertFilePath, _ := cfg.String("ssl.cert")

	//TODO factor out the for loops into separate functions, this is very bad repetetive code...
	headers, _ := cfg.List("cors.headers")
	sheaders := make([]string, len(headers))
	for i, v := range headers {
		sheaders[i] = fmt.Sprint(v)
	}
	origins, _ := cfg.List("cors.origins")
	sorigins := make([]string, len(origins))
	for i, v := range origins {
		sorigins[i] = fmt.Sprint(v)
	}
	methods, _ := cfg.List("cors.methods")
	smethods := make([]string, len(methods))
	for i, v := range methods {
		smethods[i] = fmt.Sprint(v)
	}

	//prepare CORS setup
	allowedHeaders := handlers.AllowedHeaders(sheaders)
	allowedOrigins := handlers.AllowedOrigins(sorigins)
	allowedMethods := handlers.AllowedMethods(smethods)

	log.Printf("MaaS server started at " + addr + ":" + port)
	return http.ListenAndServeTLS(addr+":"+port, sslServerCertFilePath, sslServerKeyFilePath, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
}

func setupSignalHandling(cfg *config.Config) {
	shutdownGraceperiodSeconds, _ := cfg.Int("general.shutdown_graceperiod_seconds")
	//TODO / nice to have: after "stop" signals are received, block processing of any further requests to the server
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("Caught SIG: %+v\n", sig)
		log.Printf("Wait for %d second(s) to finish processing\n", shutdownGraceperiodSeconds)
		time.Sleep(time.Duration(shutdownGraceperiodSeconds) * time.Second)
		os.Exit(0)
	}()
}
