package usecase

import (
	"database/sql"
	"fmt"

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
	// ins, err := db.Query("insert into room values ( 1, '6b2dfe65-22c2-4cdc-ad73-ef09b0e1ca79', 'Awesome Golang', default, default, default )")
	// if err != nil {
	// 	return err
	// }
	// defer ins.Close()

	ins, err := db.Prepare("insert into room values (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	fmt.Println("uscs_ins: ", ins)
	// ins.Exec(id, uuid, title, createdAt, updatedAt, deletedAt)
	fmt.Println("uscs_err: ", err)

	exec, err := ins.Exec(raw.ID, raw.UUID, raw.Title, raw.CreatedAt, raw.UpdatedAt, raw.DeletedAt)
	fmt.Println(err)
	fmt.Println(exec)

	return nil
}
