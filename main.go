package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Fact struct {
    gorm.Model
    Text 	string
}


// fun main()
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/facts", GetFacts).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetFacts(w http.ResponseWriter, r *http.Request) {
	// db connection
	db, err := gorm.Open("postgres", "user=rohan dbname=fun_facts sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	facts := []Fact{}
	db.Find(&facts)

	json.NewEncoder(w).Encode(facts)
}
// func GetPerson(w http.ResponseWriter, r *http.Request) {}
// func CreatePerson(w http.ResponseWriter, r *http.Request) {}
// func DeletePerson(w http.ResponseWriter, r *http.Request) {}