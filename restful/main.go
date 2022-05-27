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
	router.HandleFunc("/orders/add/{id}", controllers.AddItem).Methods("PUT")
	router.HandleFunc("/orders/remove/{id}", controllers.RemoveItem).Methods("PUT")
	router.HandleFunc("/orders/clear/{id}", controllers.ClearOrder).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
