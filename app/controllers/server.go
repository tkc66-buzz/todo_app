package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/tkc66-buzz/todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, fileNames ...string) {
	var files []string
	for _, fn := range fileNames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", fn))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
