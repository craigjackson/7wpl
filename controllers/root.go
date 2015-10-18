package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var rootTemplate *template.Template

func init() {
	dir, _ := os.Getwd() // gives us the source path if we haven't installed.
	templatesPath := filepath.Join(dir, "templates")
	var err error
	rootTemplate, err = template.ParseFiles(filepath.Join(templatesPath, "root.html"))
	if err != nil {
		log.Fatalf("Parse playersTemplate: %s\n", err.Error())
	}
}

func GetRootHTML(w http.ResponseWriter, req *http.Request) {
	err := rootTemplate.Execute(w, struct{}{})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
}
