package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	//	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/order", orderHandler)
	http.ListenAndServe(":8082", nil)
}

type orderStruct struct {
	ProductID        string `json:"ProductID"`
	Email            string `json:"Email"`
	Address          string `json:"Address"`
	Zip              string `json:"Zip"`
	CreditCardNumber string `json:"CreditCardNumber"`
}

func orderHandler(rw http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	body := string(bodyBytes[:])
	log.Println("Raw http request body ", body)

	var order orderStruct
	err = json.Unmarshal(bodyBytes, &order)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	log.Println("JSON ", order)

	db := dbConn()
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test.order VALUES ( ?,?,? )", order.ProductID, order.Email, order.CreditCardNumber)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
}
