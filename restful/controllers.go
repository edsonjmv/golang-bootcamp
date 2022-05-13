package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

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
	// TODO item must be received in the request body
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
	// TODO item must be received in the request body
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
