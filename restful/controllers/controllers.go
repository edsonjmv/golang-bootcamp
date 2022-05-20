package controllers

import (
	"encoding/json"
	"net/http"
	"restful/global"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(global.Db)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	o := make(global.Order)
	id := uuid.New()
	global.Db.Update(id, o)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]

	oid, _ := uuid.Parse(id)

	o := global.Db[oid]

	json.NewEncoder(w).Encode(o)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	// TODO item must be received in the request body
	item := params["item"]

	oid, _ := uuid.Parse(id)

	o := global.Db[oid]

	o.AddItem(item)

	json.NewEncoder(w).Encode(o)
}

func RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	// TODO item must be received in the request body
	item := params["item"]

	oid, _ := uuid.Parse(id)

	o := global.Db[oid]

	o.RemoveItem(item)

	json.NewEncoder(w).Encode(o)
}

func ClearOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]

	oid, _ := uuid.Parse(id)

	o := global.Db[oid]

	o.ClearOrder()

	json.NewEncoder(w).Encode(o)
}
