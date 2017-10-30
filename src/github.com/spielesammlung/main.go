package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/byuoitav/hateoas"
	"github.com/gorilla/mux"
	"github.com/spielesammlung/games"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/NilsEckelt/Spielesammlung/master/src/github.com/spielesammlung/swagger.yml")
	if err != nil {
		log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/gc", hateoas.RootResponse).Methods("GET")
	router.HandleFunc("/games", games.RetrieveGames).Methods("GET")
	router.HandleFunc("/games/{id}", games.GetGame).Methods("GET")
	router.HandleFunc("/games/{id}", games.CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", games.DeleteGame).Methods("DELETE")
	port := os.Getenv("PORT")
	fmt.Printf("Starting Spielesammlung at Port: %s", port)
	http.ListenAndServe(":"+port, router)
}
