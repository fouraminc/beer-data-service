package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Logger *log.Logger
}

func (a *App) InitializeDB(user, password, dbname, host, port, sslmode string) {

	a.Logger = log.New(os.Stdout, "", log.LstdFlags)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, password, dbname, host, port, sslmode)

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

