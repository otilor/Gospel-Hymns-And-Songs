package ghs


import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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