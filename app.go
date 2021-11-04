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

	a.Router.HandleFunc("/health", a.healthStatus).Methods("GET")
	a.Router.HandleFunc("/beers", a.getBeers).Methods("GET")
	a.Router.HandleFunc("/beer/{id:[0-9]+}", a.getBeer).Methods("GET")
	//a.Router.HandleFunc("/beer{id:[0-9]+}", a.getBeer).Methods("GET")
	a.Router.HandleFunc("/beer", a.createBeer).Methods("POST")

}

func (a *App) Run(addr string) {
	fmt.Println("I am here")
	loggerRouter := a.createLoggingRouter(a.Logger.Writer())

	a.Logger.Fatal(http.ListenAndServe(addr, loggerRouter))

}
