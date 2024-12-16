package db

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Import the pq package for PostgreSQL driver
)

var DB *gorm.DB

func ConnectDb() {

	dsn := "user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	if DB == nil {
		log.Println(errors.New("database connection is not initialized"))
	}

}
