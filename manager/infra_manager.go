package manager

import (
	"database/sql"
	"log"
	"sync"
	
	_ "github.com/lib/pq"
)

// gunanya untuk koneksi ke database
type InfraManager interface {
	GetDB() *sql.DB
}

type infraManager struct {
	db *sql.DB 
}

var onceLoadDB sync.Once

func (im *infraManager) GetDB() *sql.DB {
	// .DO hanyak akan dijalankan sekali saja  
	onceLoadDB.Do(func() {
	db, err := sql.Open("postgres", "user=postgres host=localhost password=adelia0609 dbname=laundry sslmode=disable")
	if err != nil {
			log.Fatal("Cannot start app, error when connect to DB", err.Error())
	}
	im.db = db
	})
	return im.db
}

func NewInfraManager() InfraManager {
	// objek struct
	return &infraManager{}
}