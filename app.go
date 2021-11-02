package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {

	Router *mux.Router
	Logger *log.Logger
}

func (a *App) Initialize() {

	a.Logger = log.New(os.Stdout, "",log.LstdFlags)
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/health",a.healthStatus).Methods("GET")

}

func (a *App) Run(addr string) {
	fmt.Println("I am here")
	loggerRouter := a.createLoggingRouter(a.Logger.Writer())

	a.Logger.Fatal(http.ListenAndServe(addr, loggerRouter))

}

func (a *App) healthStatus(writer http.ResponseWriter, request *http.Request) {

	response, _ := json.Marshal(struct{Status string`json:"status"`}{"OK"})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
	
}
