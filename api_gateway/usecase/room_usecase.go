package usecase

import (
	"database/sql"
	"os"

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
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open(driver, user+":"+password+"@"+"("+host+")"+"/"+dbname)
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
