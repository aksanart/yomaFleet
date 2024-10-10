package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/aksanart/weplus/dashboard/router"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	http.HandleFunc("/", router.Home)
	http.HandleFunc("/register", router.DisplayRegister)
	http.HandleFunc("/create-vehicle", router.DisplayAddVehicle)
	http.HandleFunc("/submit-vehicle", router.AddVehicle)
	http.HandleFunc("/vehicle-edit", router.DisplayEditVehicle)
	http.HandleFunc("/vehicle-delete", router.DeleteVehicle)
	http.HandleFunc("/vehicle-edit-submit", router.SubmitEditVehicle)
	http.HandleFunc("/post-register", router.SubmitRegister)
	http.HandleFunc("/vehicle-track", router.LiveTracking)
	http.HandleFunc("/login", router.LoginHandler)
	http.HandleFunc("/list-vehicle", router.ListVehicle)
	http.HandleFunc("/vehicle-detail", router.DetailVehicle)
	go func() {
		url := "http://localhost:1234"
		openBrowser(url)
	}()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal(err)
	}
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		panic(err)
	}
}
