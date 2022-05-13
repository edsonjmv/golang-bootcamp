package main

import "github.com/google/uuid"

func (d *database) update(id uuid.UUID, order order) {
	(*d)[id] = order
}
