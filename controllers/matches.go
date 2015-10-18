package controllers

import (
	"7wpl/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var matchesTemplate *template.Template
var matchTemplate *template.Template

func init() {
	dir, _ := os.Getwd() // gives us the source path if we haven't installed.
	templatesPath := filepath.Join(dir, "templates")
	var err error
	matchesTemplate, err = template.ParseFiles(filepath.Join(templatesPath, "matches.html"))
	if err != nil {
		log.Fatalf("Parse matchesTemplate: %s\n", err.Error())
	}
	matchTemplate, err = template.ParseFiles(filepath.Join(templatesPath, "match.html"))
	if err != nil {
		log.Fatalf("Parse matchTemplate: %s\n", err.Error())
	}
}

func getMatches(req *http.Request) ([]*models.Match, error) {
	matches, err := models.GetMatches()
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetMatchesHTML(w http.ResponseWriter, req *http.Request) {
	matches, err := getMatches(req)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	err = matchesTemplate.Execute(w, struct{ Matches []*models.Match }{Matches: matches})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
}

func CreateMatch(w http.ResponseWriter, req *http.Request) {
	dateStr := req.FormValue("date")
	if dateStr == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "You forgot to give a date")
		return
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "The date needs to be formatted like YYYY-MM-DD")
		return
	}
	match := &models.Match{Date: date}
	err = match.Save()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/matches/%d", match.Id))
	w.WriteHeader(302)
	fmt.Fprint(w, "Redirecting...")
}

func getMatch(req *http.Request) (*models.Match, error) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	matches, err := models.GetMatch(id)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetMatchHTML(w http.ResponseWriter, req *http.Request) {
	match, err := getMatch(req)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	err = matchTemplate.Execute(w, match)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
}
