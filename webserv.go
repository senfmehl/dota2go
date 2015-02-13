package main

import (
	"./dota2go/src/players"
	_ "database/sql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

// Handler for the MainRoute!
// Display Introduction Webpage with links to TeamStats & PlayerStats
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Rendering of Template home worked fine")
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]
	//log.Printf(player)
	err := templates.ExecuteTemplate(w, "player", player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Rendering of Template player worked fine")
}

func addPlayerHandler(w http.ResponseWriter, r *http.Request) {

	// Renders a Form to add a new Player to the playerlist
	// Stores Data input
}

func main() {
	var p players.Player
	/*
		r := mux.NewRouter()
		r.HandleFunc("/", homeHandler)
		r.HandleFunc("/players/", addPlayerHandler).Methods("POST")
		r.HandleFunc("/players/{player}", playersHandler)
		http.Handle("/", r)
		http.ListenAndServe(":8080", nil)
	*/
}
