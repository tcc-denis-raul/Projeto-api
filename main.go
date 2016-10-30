package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tcc-denis-raul/Projeto-api/api"
)

var DefaultPort = "5000"

func main() {

	router := api.NewRouter()
	port := DefaultPort
	if os.Getenv("API_PORT") != "" {
		port = os.Getenv("API_PORT")
	}
	log.Println("listening on port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
