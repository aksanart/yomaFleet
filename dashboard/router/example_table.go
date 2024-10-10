package router

import (
	"fmt"
	"html/template"
	"net/http"
)

func DisplayExample(w http.ResponseWriter, r *http.Request) {
	html := "./html/example_table.htm"
	headerTemplate := includeHeader(html)
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Sprintln(cookie.Value)
	data := PageData{
		Title: "Home Page",
	}
	err = headerTemplate.ExecuteTemplate(w, "example_table.htm", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func includeHeader(html string) *template.Template {
	return template.Must(template.ParseFiles("./html/header.htm", html))
}
