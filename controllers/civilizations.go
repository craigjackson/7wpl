package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"7wpl/models"
	"github.com/gorilla/mux"
)

func GetCivilizations(w http.ResponseWriter, req *http.Request) {
	body, err := json.Marshal(models.GetAllCivilizations())
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(body))
}

func GetCivilization(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad ID")
		return
	}
	civilization := models.GetCivilization(id)
	if civilization == nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404 page not found")
		return
	}
	body, err := json.Marshal(civilization)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(body))
}
