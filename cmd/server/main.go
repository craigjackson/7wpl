package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"7wpl/controllers"
	"github.com/gorilla/mux"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
}

func main() {
	r := mux.NewRouter()
	apiRoutes(r.PathPrefix("/api").Subrouter())
	r.Path("/players").Methods("GET").HandlerFunc(controllers.GetPlayersHTML)
	r.Path("/players").Methods("POST").HandlerFunc(controllers.CreatePlayer)
	r.Path("/players/{id:[0-9]+}").Methods("GET").HandlerFunc(controllers.GetPlayerHTML)
	r.Path("/matches").Methods("GET").HandlerFunc(controllers.GetMatchesHTML)
	r.Path("/matches").Methods("POST").HandlerFunc(controllers.CreateMatch)
	r.Path("/matches/{id:[0-9]+}").Methods("GET").HandlerFunc(controllers.GetMatchHTML)
	r.Path("/").Methods("GET").HandlerFunc(controllers.GetRootHTML)
	http.Handle("/", r)
	server := fmt.Sprintf(":%s", port)
	log.Printf("Listening on %s...", server)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Printf("ListenAndServe: %s\n", err.Error())
	}
}

func apiRoutes(r *mux.Router) {
	r.Path("/players").Methods("GET").HandlerFunc(controllers.GetPlayersJSON)
	r.Path("/players/{id:[0-9]+}").Methods("GET").HandlerFunc(controllers.GetPlayerJSON)
	r.Path("/civilizations").Methods("GET").HandlerFunc(controllers.GetCivilizations)
	r.Path("/civilizations/{id:[0-9]+}").Methods("GET").HandlerFunc(controllers.GetCivilization)
}
