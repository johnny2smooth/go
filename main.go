package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Content struct {
	Title string
	Body  string
}

var content = []Content{{Title: "hiyo", Body: "funker"}, {Title: "miso", Body: "funky"}}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Errorf("Ahhh!")
	}
	html.Execute(w, content)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", handleHome)

	http.ListenAndServe("localhost:4321", server)
}
