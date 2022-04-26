package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n2, n1 := 0, 1
	return func() int {
		n := n2
		n2, n1 = n1, n+n1
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
