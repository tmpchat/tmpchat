package controller

import (
	"net/http"
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
}

func (rc roomController) List(w http.ResponseWriter, r *http.Request) {
}

func (rc roomController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
}
