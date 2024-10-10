package router

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/aksanart/weplus/dashboard/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type PageData struct {
	Title   string
	Header  string
	Content string
}

func LiveTracking(w http.ResponseWriter, r *http.Request) {
	html := "./html/vehicle_live_tracking.htm"
	tmpl := includeHeader(html)
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	conn, err := grpc.NewClient(os.Getenv("GRPC_STREAM"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := contract.NewApiGatewayClient(conn)
    md := metadata.New(map[string]string{"session-id": cookie.Value})
    ctx := metadata.NewOutgoingContext(context.Background(), md)
	id := r.URL.Query().Get("id")
	stream, err := client.LiveTrackingOne(ctx, &contract.LiveTrackingOneReq{
		VehicleId: id,
	})
	if err != nil {
		log.Fatalf("could not start stream: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			tmpl, err := template.ParseFiles("./html/error.htm")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_ = tmpl.Execute(w, response)
			return
		}
		response.Title="Live Tracking Locaton"
		err = tmpl.ExecuteTemplate(w, "vehicle_live_tracking.htm", response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
}