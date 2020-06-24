package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/order", order)
	http.ListenAndServe(":8082", nil)
}

type orderStruct struct {
	productID        string
	address          string
	zip              string
	creditCardNumber string
}

func order(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var o orderStruct
	err := decoder.Decode(&o)
	if err != nil {
		panic(err)
	}
	log.Println(o.productID)
}
