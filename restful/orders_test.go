package main

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	o := make(order)

	o.addItem("Bananas")
	o.addItem("Apple")
	o.addItem("Pear")

	if len(o) != 3 {
		t.Errorf("Expected order length of 3, but got %v", len(o))
	}

	if o["Bananas"] == 0 {
		t.Errorf("Item expected to be in order")
	}

	if o["Apple"] == 0 {
		t.Errorf("Item expected to be in order")
	}

	if o["Pear"] == 0 {
		t.Errorf("Item expected to be in order")
	}
}

func TestRemoveItem(t *testing.T) {
	o := make(order)

	o.addItem("Bananas")
	o.addItem("Apple")
	o.addItem("Pear")

	if o["Apple"] == 0 {
		t.Errorf("Item expected to be in order")
	}

	o.removeItem("Apple")

	if o["Apple"] != 0 {
		t.Errorf("Item expected to be removed from order")
	}
}

func TestClear(t *testing.T) {
	o := make(order)

	o.addItem("Bananas")
	o.addItem("Apple")
	o.addItem("Pear")

	if len(o) != 3 {
		t.Errorf("Expected order length of 3, but got %v", len(o))
	}

	o.clearOrder()

	if len(o) != 0 {
		t.Errorf("Expected order length of 0, but got %v", len(o))
	}
}
