package controllers

import (
	"7wpl/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

var playersTemplate *template.Template
var playerTemplate *template.Template

func init() {
	dir, _ := os.Getwd() // gives us the source path if we haven't installed.
	templatesPath := filepath.Join(dir, "templates")
	var err error
	playersTemplate, err = template.ParseFiles(filepath.Join(templatesPath, "players.html"))
	if err != nil {
		log.Fatalf("Parse playersTemplate: %s\n", err.Error())
	}
	playerTemplate, err = template.ParseFiles(filepath.Join(templatesPath, "player.html"))
	if err != nil {
		log.Fatalf("Parse playerTemplate: %s\n", err.Error())
	}
}

func getPlayers(req *http.Request) ([]*models.Player, error) {
	players, err := models.GetPlayers()
	if err != nil {
		return nil, err
	}
	return players, nil
}

func GetPlayersHTML(w http.ResponseWriter, req *http.Request) {
	players, err := getPlayers(req)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	err = playersTemplate.Execute(w, struct{ Players []*models.Player }{Players: players})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
}

func GetPlayersJSON(w http.ResponseWriter, req *http.Request) {
	players, err := getPlayers(req)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	body, err := json.Marshal(players)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(body))
}

func CreatePlayer(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	if name == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "You forgot to give a name")
		return
	}
	player := &models.Player{Name: name}
	err := player.Save()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/players/%d", player.Id))
	w.WriteHeader(302)
	fmt.Fprint(w, "Redirecting...")
}

func getPlayer(req *http.Request) (*models.Player, error) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	player, err := models.GetPlayer(id)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func GetPlayerHTML(w http.ResponseWriter, req *http.Request) {
	player, err := getPlayer(req)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			fmt.Fprint(w, "404 page not found")
			return
		}
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	err = playerTemplate.Execute(w, player)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
}

func GetPlayerJSON(w http.ResponseWriter, req *http.Request) {
	player, err := getPlayer(req)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			fmt.Fprint(w, "404 page not found")
			return
		}
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	body, err := json.Marshal(player)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(body))
}
