package main

import "fmt"

func Calculate(operation string, x int, y int) {
	var operator string
	result := 0
	switch operation {
	case "add":
		operator = "+"
		result = add(x, y)
	case "subtract":
		operator = "-"
		result = subtract(x, y)
	case "multiply":
		operator = "*"
		result = multiply(x, y)
	case "divide":
		operator = "/"
		result = divide(x, y)
	default:
		fmt.Println("Invalid Operation")
	}

	fmt.Printf("%d %s %d = %d", x, operator, y, result)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}
