package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with DB
func New(d Driver) {
	switch d {
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error

		dsn := "host=localhost user=postgres password=E5p1n0z4% dbname=go_db_gorm sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("can´t open db: %v", err)
		}

		fmt.Println("Connected to Postgres")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
