package main

import (
	//_ "database/sql"
	"github.com/gorilla/mux"
	"github.com/senfmehl/dota2go/app"
	"html/template"
	"log"
	"net/http"
	_ "reflect"
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
	// fetch player from db
	var p app.Player
	p.GetPlayer(vars["player"])
	err := templates.ExecuteTemplate(w, "player", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Rendering of Template player worked fine")
}

func newPlayerHandler(w http.ResponseWriter, r *http.Request) {
	var p app.Player
	err := templates.ExecuteTemplate(w, "newPlayer", player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renders a Form to add a new Player to the playerlist
	// Stores Data input
}

func savePlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

}

func main() {
	//p := app.Player{Id: 1, Name: "Thees"}
	//log.Println(reflect.TypeOf(p))
	//log.Println(p)

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/players/", newPlayerHandler)
	r.HandleFunc("/players/{player}", playersHandler)
	r.HandleFunc("/players/save/{player}", savePlayerHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
