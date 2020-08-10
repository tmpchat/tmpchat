package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/tmpchat/tmpchat/api_gateway/controller"
)

func main() {
	// routing
	router := mux.NewRouter()
	con := controller.NewRoomController()
	router.HandleFunc("/room/create", con.Create)
	router.HandleFunc("/room/find", con.Find)
	router.HandleFunc("/room/list", con.List)
	router.HandleFunc("/room/update-title", con.UpdateTitle)

	// listen
	apiGatewayHost := os.Getenv("API_GATEWAY_HOST")
	if apiGatewayHost == "" {
		apiGatewayHost = "127.0.0.1:8888"
	}
	log.Fatal(http.ListenAndServe(apiGatewayHost, router))

	fmt.Println("Done")
}
