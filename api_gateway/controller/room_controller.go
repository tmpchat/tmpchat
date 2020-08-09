package controller

import (
	"fmt"
	"net/http"

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
	uscs := usecase.NewRoomUsecase()
	uuid := "9f0c3721-cc15-451c-8700-d5d5af0c677f"
	row, err := uscs.Find(uuid)
	if err != nil {
		fmt.Println("err: ", err)
		// TODO: HTTP response 400
	}
	// TODO: Response RoomEntity to Client
	fmt.Printf(`RoomController.List: %#v, %#v`, w, r)
	fmt.Println("RoomEntity: ", row)
}

func (rc roomController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.List: %#v, %#v`, w, r)
}

func (rc roomController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`RoomController.UpdateTitle: %#v, %#v`, w, r)
}
