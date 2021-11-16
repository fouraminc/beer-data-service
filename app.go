package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Logger *log.Logger
	Config
}

type Config struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
	sslMode    string
}

func (a *App) InitializeConfig() {
	config := Config{}

	vi := viper.New()
	vi.SetConfigFile("config.yaml")
	vi.ReadInConfig()

	config.dbUser = vi.GetString("dbUser")
	config.dbPassword = vi.GetString("dbPassword")
	config.sslMode = vi.GetString("sslMode")
	config.dbName = vi.GetString("dbName")
	config.dbHost = vi.GetString("dbHost")
	config.dbPort = vi.GetString("dbPort")
	a.Config = config

}

func (a *App) InitializeDB() {

	fmt.Println("App Config:")
	fmt.Println(a.Config)
	a.Logger = log.New(os.Stdout, "", log.LstdFlags)

	fmt.Println(a.Config.dbHost)

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		a.Config.dbUser,
		a.Config.dbPassword,
		a.Config.dbHost,
		a.Config.dbPort,
		a.Config.dbName,
		a.Config.sslMode)
	a.Logger.Println(dsn)
	var err error
	a.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		a.Logger.Fatal(err)
	}

}

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()

	// Am I up
	a.Router.HandleFunc("/health", a.healthStatus).Methods("GET")

	// Get all beers
	a.Router.HandleFunc("/beers", a.getBeers).Methods("GET")

	// handle path variables
	a.Router.HandleFunc("/beer/{id:[0-9]+}", a.getBeer).Methods("GET")

	// handle query parameters

	a.Router.HandleFunc("/beer", a.getBeer).Methods("GET").Queries("id", "{id:[0-9]+}")

	// beer me
	a.Router.HandleFunc("/beer", a.createBeer).Methods("POST")

	// remove beer, how sad, path variables
	a.Router.HandleFunc("/beer/{id:[0-9]+}", a.deleteBeer).Methods("DELETE")

	// remove beer, still sad, query variables
	a.Router.HandleFunc("/beer", a.deleteBeer).Methods("DELETE").Queries("id", "{id:[0-9]+}")

	// update the beer, everyone likes a new twist
	a.Router.HandleFunc("/beer/{id:[0-9]+}", a.updateBeer).Methods("PUT")

}

func (a *App) Run(addr string) {
	fmt.Println("I am here")
	loggerRouter := a.createLoggingRouter(a.Logger.Writer())

	a.Logger.Fatal(http.ListenAndServe(addr, loggerRouter))

}
