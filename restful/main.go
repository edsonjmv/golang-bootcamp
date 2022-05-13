package main

import (
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/orders/add/{id}/{item}", AddItem).Methods("PUT")
	router.HandleFunc("/orders/remove/{id}/{item}", RemoveItem).Methods("PUT")
	router.HandleFunc("/orders/clear/{id}", ClearOrder).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))

	fmt.Println(db)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	o := make(order)
	id := uuid.New()
	db.update(id, o)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]

	oid, _ := uuid.Parse(id)

	o := db[oid]

	json.NewEncoder(w).Encode(o)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	item := params["item"]

	oid, _ := uuid.Parse(id)

	o := db[oid]

	o.addItem(item)

	json.NewEncoder(w).Encode(o)
}

func RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	item := params["item"]

	oid, _ := uuid.Parse(id)

	o := db[oid]

	o.removeItem(item)

	json.NewEncoder(w).Encode(o)
}

func ClearOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]

	oid, _ := uuid.Parse(id)

	o := db[oid]

	o.clearOrder()

	json.NewEncoder(w).Encode(o)
}

func (d *database) update(id uuid.UUID, order order) {
	(*d)[id] = order
}

func (o *order) addItem(item string) {
	if (*o)[item] > 0 {
		(*o)[item]++
	} else {
		(*o)[item] = 1
	}
}

func (o *order) removeItem(item string) {
	delete(*o, item)
}

func (o *order) clearOrder() {
	for item := range *o {
		delete(*o, item)
	}
}
