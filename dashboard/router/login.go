package router

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aksanart/weplus/dashboard/config"
	"github.com/aksanart/weplus/dashboard/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	payload := model.LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	res, err := config.PostHttp(payload, "login", []config.Headers{})
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
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   loginRes.Data.SessionId,
			Expires: time.Now().Add(150 * time.Minute),
		})
		http.Redirect(w, r, "/list-vehicle", http.StatusSeeOther)
	}
}
