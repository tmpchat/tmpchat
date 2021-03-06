package gateway

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/tmpchat/tmpchat/api_gateway/domain"
)

type RoomRepository interface {
	DBConn() (db *sql.DB)
	Create(room domain.CreateRoomRequest) (*domain.RoomEntity, error)
	Find(id string) (*domain.RoomEntity, error)
	List() ([]*domain.RoomEntity, error)
	UpdateTitle(req domain.UpdateTitleRequest) (*domain.RoomEntity, error)
	Delete(req domain.DeleteRoomRequest) error
}

// TODO: Isn't bad performance to call DBConn every time the method is called?
type roomRepository struct{}

func NewRoomRepository() RoomRepository {
	return roomRepository{}
}

func (r roomRepository) DBConn() (db *sql.DB) {
	driver := "mysql"
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open(driver, user+":"+password+"@"+"("+host+")"+"/"+dbname+"?parseTime=true")
	if err != nil {
		// TODO: remove panic
		panic(err.Error())
	}
	return db
}

func (r roomRepository) Create(room domain.CreateRoomRequest) (*domain.RoomEntity, error) {
	db := r.DBConn()
	_, err := db.Query("insert into room values (default, ?, ?, default, default, ?)", room.UUID, room.Title, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := r.Find(room.UUID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r roomRepository) Find(id string) (*domain.RoomEntity, error) {
	var row domain.RoomEntity
	db := r.DBConn()
	err := db.QueryRow("select * from room where external_id = ? and deleted_at is null", id).Scan(&row.ID, &row.UUID, &row.Title, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &row, nil
}

func (r roomRepository) List() ([]*domain.RoomEntity, error) {
	db := r.DBConn()
	// Execute the query
	rows, err := db.Query("select * from room where deleted_at is null")
	if err != nil {
		return nil, err
	}

	var rooms []*domain.RoomEntity
	// Fetch rows
	for rows.Next() {
		var room domain.RoomEntity
		err = rows.Scan(&room.ID, &room.UUID, &room.Title, &room.CreatedAt, &room.UpdatedAt, &room.DeletedAt)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, &room)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r roomRepository) UpdateTitle(req domain.UpdateTitleRequest) (*domain.RoomEntity, error) {
	db := r.DBConn()
	_, err := db.Query("update room set title = ? where external_id = ? and deleted_at is null", req.Title, req.UUID)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := r.Find(req.UUID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r roomRepository) Delete(req domain.DeleteRoomRequest) error {
	db := r.DBConn()
	_, err := db.Query("update room set deleted_at = ? where external_id = ? and deleted_at is null", time.Now(), req.UUID)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}
