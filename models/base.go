package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load() //load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbHost, dbName)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	// db.Debug().AutoMigrate(&Account{}, &Contact{})
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}