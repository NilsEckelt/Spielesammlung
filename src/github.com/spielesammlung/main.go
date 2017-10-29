package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spielesammlung/games"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/games", games.RetrieveGames).Methods("GET")
	router.HandleFunc("/games/{id}", games.GetGame).Methods("GET")
	router.HandleFunc("/games/{id}", games.CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", games.DeleteGame).Methods("DELETE")
	port := os.Getenv("PORT")
	fmt.Printf("Starting Spielesammlung at Port: %s", port)
	http.ListenAndServe(":"+port, router)
}
