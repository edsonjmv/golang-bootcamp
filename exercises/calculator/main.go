package main

import "fmt"

func main() {
	var operation string
	var n1, n2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&n1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&n2)
	fmt.Print("Enter a operation (add, subtract, multiply, divide):")
	fmt.Scanln(&operation)

	Calculate(operation, n1, n2)
}
