package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tmpchat/tmpchat/api_gateway/controller"
)

func  main()  {
	// routing
	con := controller.NewRoomController()
	http.HandleFunc("/room/create", con.Create)
	http.HandleFunc("/room/list", con.List)
	http.HandleFunc("/room/update-title", con.UpdateTitle)

	// listen
	log.Fatal(http.ListenAndServe(":8888", nil))

	fmt.Println("Done")
}
