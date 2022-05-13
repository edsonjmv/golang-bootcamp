package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type order map[string]int

type database map[uuid.UUID]order

var db = make(database)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/orders", CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", GetOrder).Methods("GET")
	// TODO item must be received in the request body
	router.HandleFunc("/orders/add/{id}/{item}", AddItem).Methods("PUT")
	// TODO item must be received in the request body
	router.HandleFunc("/orders/remove/{id}/{item}", RemoveItem).Methods("PUT")
	router.HandleFunc("/orders/clear/{id}", ClearOrder).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
