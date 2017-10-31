package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	db "github.com/nilseckelt/spielesammlung/database"
)

func getHostFromUrl(r *http.Request) string {
	if strings.HasPrefix(r.Host, "localhost") {
		return "http://" + r.Host
	}
	return r.URL.Scheme + r.Host
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	database := db.New("localhost")
	defer database.Close()

	if r.Method == "GET" {
		games := database.RetrieveGames()
		for i, game := range games {
			game.AddHosts(getHostFromUrl(r))
			games[i] = game
		}
		json.NewEncoder(w).Encode(games)
	}
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	database := db.New("localhost")
	if r.Method == "GET" {
		game, err := database.Find(vars["gameName"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			game.AddHosts(getHostFromUrl(r))
			json.NewEncoder(w).Encode(game)
		}
	}
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func EditGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
