package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// send a payload of JSON content
func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// send a JSON error message
func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})

	a.Logger.Printf("App error: code %d, message %s", code, message)
}

func (a *App) createBeer(w http.ResponseWriter, r *http.Request) {
	var b beer
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&b); err != nil {

		msg := fmt.Sprintf("Invaled request payload. Error: %s", err.Error())
		a.respondWithError(w, http.StatusBadRequest, msg)
	}
	defer r.Body.Close()

	if err := b.createBeer(a.DB); err != nil {
		a.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondWithJSON(w, http.StatusCreated, b)
}

func (a *App) healthStatus(writer http.ResponseWriter, request *http.Request) {

	response, _ := json.Marshal(struct {
		Status string `json:"status"`
	}{"OK"})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)

}

func (a *App) getBeers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In getBeers")

	beers, err := getBeers(a.DB)

	if err != nil {
		a.respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	a.respondWithJSON(w, http.StatusOK, beers)
}

func (a *App) getBeer(w http.ResponseWriter, r *http.Request) {

	fmt.Println("In getBeer")
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	if err != nil {
		msg := fmt.Sprintf("Invalid id.  Error: %s", err.Error())
		a.respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	b := beer{ID: id}
	if err := b.getBeer(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			msg := fmt.Sprintf("Beer not found.")
			a.respondWithError(w, http.StatusNotFound, msg)
		default:
			a.respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	a.respondWithJSON(w, http.StatusOK, b)
}


func (a *App) deleteBeer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In deleteBeer")

	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	if err != nil {
		a.respondWithError(w, http.StatusBadRequest, "Invalid argument" + err.Error())
	}

	b := beer{ID :id}

	if err := b.deleteBeer(a.DB); err != nil {
		fmt.Println("Maybe and error" + err.Error())
		switch err {

		// maybe overkill here?
		case sql.ErrNoRows:
			a.respondWithError(w, http.StatusNotFound, "Beer not found")

		default:
			a.respondWithError(w, http.StatusInternalServerError, "Something went wrong"+ err.Error())
		}
		return
	}
	return
}
