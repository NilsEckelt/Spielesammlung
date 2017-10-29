package games

import (
	"encoding/json"
	"net/http"

	"github.com/spielesammlung/cupboard"
)

func RetrieveGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cupboard := cupboard.New("localhost")
	defer cupboard.Close()

	if r.Method == "GET" {
		games := cupboard.RetrieveGames()
		json.NewEncoder(w).Encode(games)
	}
}
func GetGame(w http.ResponseWriter, r *http.Request)    {}
func CreateGame(w http.ResponseWriter, r *http.Request) {}
func DeleteGame(w http.ResponseWriter, r *http.Request) {}
