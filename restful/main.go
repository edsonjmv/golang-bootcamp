package main

import (
	"log"
	"net/http"
	"restful/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", controllers.GetOrder).Methods("GET")
	// TODO item must be received in the request body
	router.HandleFunc("/orders/add/{id}/{item}", controllers.AddItem).Methods("PUT")
	// TODO item must be received in the request body
	router.HandleFunc("/orders/remove/{id}/{item}", controllers.RemoveItem).Methods("PUT")
	router.HandleFunc("/orders/clear/{id}", controllers.ClearOrder).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
