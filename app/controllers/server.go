package controllers

import (
	"net/http"

	"github.com/tkc66-buzz/todo_app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
