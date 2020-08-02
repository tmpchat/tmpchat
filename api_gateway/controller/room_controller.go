package controller

import (
	"fmt"
	"net/http"

	"github.com/tmpchat/tmpchat/api_gateway/domain"
	"github.com/tmpchat/tmpchat/api_gateway/usecase"
)

type RoomController interface {
	Create(w http.ResponseWriter, r *http.Request)
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
	fmt.Println(uuid)
	// TODO: Create MessageBroker
	// TODO: Insert to DB
	room := domain.CreateRoomRequest{UUID: "XXXXXXXX", Title: "Awesome Golang"}
	ins := uscs.InsertDB(room)
	fmt.Println("ins: ", ins)
	// TODO: Response RoomEntity to Client
	fmt.Printf(`RoomController.Create: %#v, %#v`, w, r)
}

func (rc roomController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.List: %#v, %#v`, w, r)
}

func (rc roomController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.UpdateTitle: %#v, %#v`, w, r)
}
