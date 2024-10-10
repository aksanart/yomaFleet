package router

import (
	"html/template"
	"net/http"
)

func DisplayError(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./html/error.htm")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tmpl.Execute(w, nil)
}
