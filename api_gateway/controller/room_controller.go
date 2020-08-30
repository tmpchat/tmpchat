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
	Delete(w http.ResponseWriter, r *http.Request)
}

type roomController struct{}

func NewRoomController() RoomController {
	return roomController{}
}

func (rc roomController) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Set '*' to Environment Variables??
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}
	uscs := usecase.NewRoomUsecase()

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify title", http.StatusBadRequest)
		return
	}

	request := domain.NewCreateRoomRequest()
	if err := json.Unmarshal(body, &request); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify title", http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify title", http.StatusBadRequest)
		return
	}

	response, err := uscs.Create(request)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "Please retry", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (rc roomController) Find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	uscs := usecase.NewRoomUsecase()
	rooms, err := uscs.List()
	if err != nil {
		// TODO: error handling
		http.Error(w, "failed list error", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, "failed marahl error", http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func (rc roomController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	uscs := usecase.NewRoomUsecase()

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify uuid and title", http.StatusBadRequest)
		return
	}

	request := domain.UpdateTitleRequest{}
	if err := json.Unmarshal(body, &request); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify uuid and title", http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify uuid and title", http.StatusBadRequest)
		return
	}

	response, err := uscs.UpdateTitle(request)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "Please retry", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (rc roomController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	uscs := usecase.NewRoomUsecase()

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify uuid", http.StatusBadRequest)
		return
	}

	request := domain.DeleteRoomRequest{}
	if err := json.Unmarshal(body, &request); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify uuid", http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "please specify uuid", http.StatusBadRequest)
		return
	}

	if err := uscs.Delete(request); err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "room is not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
