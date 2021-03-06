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
	router.HandleFunc("/rooms", con.Create).Methods("POST", http.MethodOptions)
	router.HandleFunc("/rooms", con.List).Methods("GET")
	router.HandleFunc("/rooms", con.Delete).Methods("DELETE")
	router.HandleFunc("/rooms/{id}", con.Find).Methods("GET")
	router.HandleFunc("/rooms/titles", con.UpdateTitle).Methods("PUT")

	router.Use(mux.CORSMethodMiddleware(router))

	// listen
	apiGatewayHost := os.Getenv("API_GATEWAY_HOST")
	if apiGatewayHost == "" {
		apiGatewayHost = "127.0.0.1:8888"
	}
	log.Fatal(http.ListenAndServe(apiGatewayHost, router))

	fmt.Println("Done")
}
