package render

import (
	"fmt"
	"html/template"
	"net/http"
	"web3/models"
	"web3/pkg/helpers"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, pageData *models.PageData) {
	var tmpl *template.Template
	var err error
	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Teplate in cache")
	}

	tmpl = tmplCache[t]
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println(err)
	}
}

func makeTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("../../templates/%s", t),
		"../../templates/base.layout.html",
	}

	tmpl, err := template.ParseFiles(templates...)

	helpers.ErrorCheck(err)

	tmplCache[t] = tmpl
	return nil
}
