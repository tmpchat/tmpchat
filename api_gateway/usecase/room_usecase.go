package usecase

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	google_uuid "github.com/google/uuid"
	"github.com/tmpchat/tmpchat/api_gateway/domain"
)

type RoomUsecase interface {
	CreateUUID() google_uuid.UUID
	InsertDB(raw domain.RoomEntity) error
}

type roomUsecase struct{}

func NewRoomUsecase() RoomUsecase {
	return roomUsecase{}
}

func (r roomUsecase) CreateUUID() google_uuid.UUID {
	return google_uuid.UUID(google_uuid.New())
}

func (r roomUsecase) InsertDB(raw domain.RoomEntity) error {
	db, err := sql.Open("mysql", "root:mypassword@tcp(0.0.0.0:3306)/tmpchat")
	if err != nil {
		return err
	}
	defer db.Close()

	ins, err := db.Query("insert into room values ( default, ?, ?, default, default, ? )", raw.UUID, raw.Title, nil)
	if err != nil {
		return err
	}
	defer ins.Close()

	return err
}
