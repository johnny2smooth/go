package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/johnny2smooth/go/data"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Errorf("Ahhh!")
	}
	html.Execute(w, data.MyCount)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/clicked", data.HandleClick)
	server.HandleFunc("/", handleHome)

	http.ListenAndServe("localhost:4321", server)
}
