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

func DisplayEditVehicle(w http.ResponseWriter, r *http.Request) {
	html := "./html/edit_vehicle.htm"
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
	headers := []config.Headers{
		{Key: "Session-Id", Value: cookie.Value},
	}
	res, err := config.GetHttp("vehicle/"+r.URL.Query().Get("id"), headers)
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
	}
	var resp model.HistoryLocationVehicleResponse
	resp.Title = "Vehicle History Location"
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		log.Fatal(err)
	}

	err = headerTemplate.ExecuteTemplate(w, "edit_vehicle.htm", resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
