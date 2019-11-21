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

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbHost, dbName)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", "postgres://zptmemqneaixmt:1e92c2cba5790b83e52506765ba9fd47ec83cf5d41a53db4420aac6a71a380c4@ec2-23-23-241-119.compute-1.amazonaws.com:5432/datl5hlol38lth")
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
