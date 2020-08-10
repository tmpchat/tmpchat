package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	uscs := usecase.NewRoomUsecase()
	uuid := uscs.CreateUUID()

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify title", http.StatusBadRequest)
		return
	}

	var request domain.CreateRoomRequest
	if err := json.Unmarshal(body, &request); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify title", http.StatusBadRequest)
		return
	}
	request.UUID = uuid.String()

	if err := uscs.Create(request); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "Please retry", http.StatusInternalServerError)
		return
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
