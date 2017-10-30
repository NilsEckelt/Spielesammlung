package main

import (
	"log"

	"github.com/byuoitav/hateoas"
	"github.com/spielesammlung/cmd/restapi"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/NilsEckelt/Spielesammlung/master/src/github.com/spielesammlung/swagger.yml")
	if err != nil {
		log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOfferAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// router := mux.NewRouter()
	// router.HandleFunc("/", hateoas.RootResponse).Methods("GET")
	// router.HandleFunc("/gc", games.RetrieveGames).Methods("GET")
	// router.HandleFunc("/gc/{id}", games.GetGame).Methods("GET")
	// router.HandleFunc("/gc/{id}", games.CreateGame).Methods("POST")
	// router.HandleFunc("/gc/{id}", games.UpdateGame).Methods("PUT")
	// router.HandleFunc("/gc/{id}", games.DeleteGame).Methods("DELETE")
	// port := os.Getenv("PORT")
	// fmt.Printf("Starting Spielesammlung at Port: %s", port)
	// http.ListenAndServe(":"+port, router)
}
