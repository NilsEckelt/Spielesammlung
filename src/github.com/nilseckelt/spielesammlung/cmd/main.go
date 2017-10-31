package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nilseckelt/spielesammlung/openapi"
)

func main() {
	port := os.Getenv("PORT")

	router := openapi.NewRouter()

	log.Printf("Starting Spielesammlung at Port:%s ...", port)
	http.ListenAndServe(":"+port, router)
}
