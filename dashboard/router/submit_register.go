package router

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/aksanart/weplus/dashboard/config"
	"github.com/aksanart/weplus/dashboard/model"
)

func SubmitRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	payload := model.LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	res, err := config.PostHttp(payload, "register", []config.Headers{})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode >= http.StatusBadRequest {
		var errResponse model.Error
		err := json.Unmarshal(resBody, &errResponse)
		if err != nil {
			log.Fatal(err)
		}
		tmpl, err := template.ParseFiles("./html/error.htm")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = tmpl.Execute(w, errResponse)
		return
	} else {
		var loginRes model.LoginResponse
		err := json.Unmarshal(resBody, &loginRes)
		if err != nil {
			log.Fatal(err)
		}
		tmpl, err := template.ParseFiles("./html/registered.htm")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = tmpl.Execute(w, payload)
	}
}