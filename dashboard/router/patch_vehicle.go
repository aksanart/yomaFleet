package router

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/aksanart/weplus/dashboard/config"
	"github.com/aksanart/weplus/dashboard/model"
)

func SubmitEditVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	mileage, err := strconv.Atoi(r.FormValue("mileage"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	payload := model.CreateVehicleReq{
		VehicleName:   r.FormValue("vehicle_name"),
		VehicleModel:  r.FormValue("vehicle_model"),
		VehicleStatus: r.FormValue("vehicle_status"),
		Mileage:       int32(mileage),
		LicenseNumber: r.FormValue("license_number"),
	}
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
	res, err := config.PatchHttp(payload, "vehicle/"+r.URL.Query().Get("id"), headers)
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
	http.Redirect(w, r, "/list-vehicle", http.StatusSeeOther)
}
