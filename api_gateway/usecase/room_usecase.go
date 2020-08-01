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
	DBConn() (db *sql.DB)
}

type roomUsecase struct{}

func NewRoomUsecase() RoomUsecase {
	return roomUsecase{}
}

func (r roomUsecase) CreateUUID() google_uuid.UUID {
	return google_uuid.UUID(google_uuid.New())
}

func (r roomUsecase) DBConn() (db *sql.DB) {
	driver := "mysql"
	user := "root"
	password := "mypassword"
	host := "0.0.0.0"
	dbname := "tmpchat"
	db, err := sql.Open(driver, user+":"+password+"@"+"("+host+")"+"/"+dbname)
	//	"root:mypassword@tcp(0.0.0.0:3306)/tmpchat")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (r roomUsecase) InsertDB(raw domain.RoomEntity) error {
	db := r.DBConn()
	ins, err := db.Query("insert into room values (default, ?, ?, default, default, ?)", raw.UUID, raw.Title, nil)
	if err != nil {
		return err
	}
	defer ins.Close()
	defer db.Close()

	return err
}
