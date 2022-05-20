package main

import (
	"restful/global"
	"testing"
)

func TestAddItem(t *testing.T) {
	o := make(global.Order)

	o.AddItem("Bananas")
	o.AddItem("Apple")
	o.AddItem("Pear")

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
	o := make(global.Order)

	o.AddItem("Bananas")
	o.AddItem("Apple")
	o.AddItem("Pear")

	if o["Apple"] == 0 {
		t.Errorf("Item expected to be in order")
	}

	o.RemoveItem("Apple")

	if o["Apple"] != 0 {
		t.Errorf("Item expected to be removed from order")
	}
}

func TestClear(t *testing.T) {
	o := make(global.Order)

	o.AddItem("Bananas")
	o.AddItem("Apple")
	o.AddItem("Pear")

	if len(o) != 3 {
		t.Errorf("Expected order length of 3, but got %v", len(o))
	}

	o.ClearOrder()

	if len(o) != 0 {
		t.Errorf("Expected order length of 0, but got %v", len(o))
	}
}
