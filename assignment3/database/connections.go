package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	// change the password when you will run code
	// also change password on docker-compose.yml
	db, err := gorm.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if !db.HasTable(&Books{}) { // check if table already exists
		db.AutoMigrate(&Books{})
		fmt.Printf("no table")
		seed(db)
	}
	return db
}

func seed(db *gorm.DB) {
	for _, book := range MyBooks() {
		db.Create(&book)
	}
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		sleep := time.Duration(1)
		for dbase == nil {
			sleep *= 2
			fmt.Printf("Database is unavailable. Wait %d sec\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
