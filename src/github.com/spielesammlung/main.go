package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spielesammlung/cupboard"
	"github.com/spielesammlung/games"
)

func main() {
	db := cupboard.New("localhost")
	defer db.Close()

	game := cupboard.Game{
		"Civilisation",
		cupboard.MinMax{
			Min: 2,
			Max: 6,
		},
		cupboard.MinMax{
			Min: 4,
			Max: 6,
		},
		"Lars",
		"Aufbau-/Strategie",
	}

	db.AddGame(game)

	http.HandleFunc("/games", games.GameServer)
	port := os.Getenv("PORT")
	fmt.Printf("Starting Spielesammlung at Port: %s", port)
	http.ListenAndServe(":"+port, nil)
}
