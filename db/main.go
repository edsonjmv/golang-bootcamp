package main

import "fmt"

type cart map[string]int

func main() {
	// Create a new cart
	userCart := createCart()

	fmt.Println("New cart created ->", userCart)

	// Adding items to a cart
	userCart.addItem("salmon")
	userCart.addItem("chicken")
	userCart.addItem("eggs")
	userCart.addItem("eggs")
	userCart.addItem("eggs")

	fmt.Println("Items added to cart ->", userCart)

	// Changing the quantity of an existent item in a cart
	userCart.setItemQuantity("salmon", 7)
	userCart.setItemQuantity("chicken", 5)

	fmt.Println("Setting salmon & chicken quantities ->", userCart)

	// Removing an item from a cart
	userCart.removeItem("eggs")

	fmt.Println("Removing eggs ->", userCart)

	fmt.Println("List all items of a specific cart:")

	// List all items of a specific cart
	userCart.listCartItems()

	// Clear a specific cart (remove all items).
	userCart.clearCart()

	fmt.Println("Clear a specific cart (remove all items) ->", userCart)
}

func createCart() cart {
	return make(cart)
}

func (c *cart) addItem(item string) {
	if (*c)[item] > 0 {
		(*c)[item]++
	} else {
		(*c)[item] = 1
	}
}

func (c *cart) listCartItems() {
	fmt.Println(c)
}

func (c *cart) setItemQuantity(item string, quantity int) {
	(*c)[item] = quantity
}

func (c *cart) removeItem(item string) {
	delete(*c, item)
}

func (c *cart) clearCart() {
	for item := range *c {
		delete(*c, item)
	}
}
