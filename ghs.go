package ghs

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var Numerals = map[int] string {
	12 : "Hey",
}


func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", nil)
	
}

func render(w http.ResponseWriter, tmpl string, r *http.Request){
	tmpl = fmt.Sprintf("form/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err.Error())
	}

	err = t.Execute(w, nil)
		if err != nil {
			log.Print("template exeuting error", err)
		}
	
}