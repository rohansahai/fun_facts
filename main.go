package main

import (
	"fmt"
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
	db, err = gorm.Open("postgres", "user=rohan dbname=fun_facts sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Fact{})

	startRouter()
}