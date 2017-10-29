package games

import (
	"encoding/json"
	"net/http"

	"github.com/spielesammlung/cupboard"
)

func GameServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cupboard := cupboard.New("localhost")
	defer cupboard.Close()

	games := cupboard.RetrieveGames()

	json.NewEncoder(w).Encode(games)
}
