package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tmpchat/tmpchat/api_gateway/controller"
)

func main() {
	// routing
	con := controller.NewRoomController()
	http.HandleFunc("/room/create", con.Create)
	http.HandleFunc("/room/find", con.Find)
	http.HandleFunc("/room/list", con.List)
	http.HandleFunc("/room/update-title", con.UpdateTitle)

	// listen
	apiGatewayHost := os.Getenv("API_GATEWAY_HOST")
	if apiGatewayHost == "" {
		apiGatewayHost = "127.0.0.1:8888"
	}
	log.Fatal(http.ListenAndServe(apiGatewayHost, nil))

	fmt.Println("Done")
}
