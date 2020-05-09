package ghs


import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type GHS struct {
	Id int
	Title string
	Chorus string
	Stanza_1 string
	Stanza_2 string
	Stanza_3 string
	Stanza_4 string
	Stanza_5 string
	Stanza_6 string
	Stanza_7 string
	Stanza_8 string
}


func render(w http.ResponseWriter, tmpl string, pageData []GHS){
	tmpl = fmt.Sprintf("form/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err.Error())
	}

	err = t.Execute(w, pageData)
		if err != nil {
			log.Print("template exeuting error", err)
		}
	
}