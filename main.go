package main

import (
	"fmt"
	"os"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Fact struct {
	ID        uint   `json:"id"`
    Text 	string   `json:"text"`
    CreatedAt time.Time `json:"createdAt"`
}


// fun main()
func main() {
	// db connection
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPwd := os.Getenv("DB_PWD")
	dbString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPwd)
	db, err = gorm.Open("postgres", dbString)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Fact{})

	startRouter()
}