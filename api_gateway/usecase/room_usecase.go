package usecase

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	google_uuid "github.com/google/uuid"
	"github.com/tmpchat/tmpchat/api_gateway/domain"
)

type RoomUsecase interface {
	CreateUUID() google_uuid.UUID
	InsertDB(*domain.RoomEntity) error
}

type roomUsecase struct{}

func NewRoomUsecase() RoomUsecase {
	return roomUsecase{}
}

func (r roomUsecase) CreateUUID() google_uuid.UUID {
	return google_uuid.UUID(google_uuid.New())
}

func (r roomUsecase) InsertDB(*domain.RoomEntity) error {
	raw := domain.RoomEntity{}
	db, err := sql.Open("mysql", "root:mypassword@tcp(0.0.0.0:3306)/tmpchat")
	if err != nil {
		return err
	}
	defer db.Close()

	ins, err := db.Prepare("insert into room(id, external_id, title, created_at, update_at, deleted_at) value (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	ins.Exec(raw)

	return nil
}
