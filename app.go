package main

import (
	"database/sql"
	"encoding/json"
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

func (a *App) InitializeDB(user, password, dbname, host, port string) {

	a.Logger = log.New(os.Stdout, "", log.LstdFlags)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", user, password, dbname, host, port)

	var err error
	a.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		a.Logger.Fatal(err)
	}

}

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/health", a.healthStatus).Methods("GET")
}

func (a *App) Run(addr string) {
	fmt.Println("I am here")
	loggerRouter := a.createLoggingRouter(a.Logger.Writer())

	a.Logger.Fatal(http.ListenAndServe(addr, loggerRouter))

}

func (a *App) healthStatus(writer http.ResponseWriter, request *http.Request) {

	response, _ := json.Marshal(struct {
		Status string `json:"status"`
	}{"OK"})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)

}
