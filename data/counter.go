package data

import (
	"fmt"
	"html/template"
	"net/http"
)

type Clicker struct {
	Count    int
	IsPlural bool
}

var MyCount = Clicker{
	Count:    0,
	IsPlural: false,
}

func HandleClick(w http.ResponseWriter, r *http.Request) {
	MyCount.Count++
	if !MyCount.IsPlural && MyCount.Count > 1 {
		MyCount.IsPlural = true
		fmt.Printf("Reassigned IsPlural to %v\n", MyCount.IsPlural)

	}
	fmt.Printf("clicked %d times\n", MyCount.Count)
	tmpl, _ := template.ParseFiles("./templates/button.html")

	tmpl.Execute(w, MyCount)
}
