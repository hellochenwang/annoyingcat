package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/order", orderHandler)
	http.ListenAndServe(":8082", nil)
}

type orderStruct struct {
	ProductID        string `json:"ProductID"`
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
}
