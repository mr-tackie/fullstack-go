package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template using the html/template package
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter, tmpl string) error {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs()
	}
}
