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
	room := domain.CreateRoomRequest{UUID: uuid.String(), Title: "Awesome Golang"}
	err := uscs.Create(room)
	if err != nil {
		fmt.Println("err: ", err)
		// TODO: HTTP response 400
	}
	// TODO: Response RoomEntity to Client
	fmt.Printf(`RoomController.Create: %#v, %#v`, w, r)
}

func (rc roomController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.List: %#v, %#v`, w, r)
}

func (rc roomController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.UpdateTitle: %#v, %#v`, w, r)
}
