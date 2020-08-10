package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tmpchat/tmpchat/api_gateway/domain"
	"github.com/tmpchat/tmpchat/api_gateway/usecase"
)

type RoomController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	UpdateTitle(w http.ResponseWriter, r *http.Request)
}

type roomController struct{}

func NewRoomController() RoomController {
	return roomController{}
}

func (rc roomController) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Recieve title
	uscs := usecase.NewRoomUsecase()
	uuid := uscs.CreateUUID()
	room := domain.CreateRoomRequest{UUID: uuid.String(), Title: "Awesome Golang"}
	err := uscs.Create(room)
	if err != nil {
		fmt.Println("err: ", err)
		// TODO: HTTP response 400
	}
	// TODO: Response RoomEntity to Client
	fmt.Printf(`RoomController.Create: %#v, %#v`, w, r)
}

func (rc roomController) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, exists := vars["id"]; !exists {
		http.Error(w, "please specify room id", http.StatusBadRequest)
		return
	}
	uuid := vars["id"]
	uscs := usecase.NewRoomUsecase()
	row, err := uscs.Find(uuid)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "room is not found", http.StatusNotFound)
		return
	}
	fmt.Printf(`RoomController.List: %#v, %#v`, w, r)

	res, err := json.Marshal(row)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func (rc roomController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.List: %#v, %#v`, w, r)
}

func (rc roomController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.UpdateTitle: %#v, %#v`, w, r)
}
