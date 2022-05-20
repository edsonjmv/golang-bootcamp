package global

import "github.com/google/uuid"

type Order map[string]int

type Database map[uuid.UUID]Order

var Db = make(Database)

func (o *Order) AddItem(item string) {
	if (*o)[item] > 0 {
		(*o)[item]++
	} else {
		(*o)[item] = 1
	}
}

func (o *Order) RemoveItem(item string) {
	delete(*o, item)
}

func (o *Order) ClearOrder() {
	for item := range *o {
		delete(*o, item)
	}
}

func (d *Database) Update(id uuid.UUID, order Order) {
	(*d)[id] = order
}
