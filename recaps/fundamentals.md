### Recap

You just read a bunch of text so lets do a set of questions, sort of like a recap, to test your knowledge. This list has not only questions but also exercises that you need to solve. We recommend you to put all this in a repository so that the instructors can see how are doing, do a bit of code review and help you with whatever difficulty you may be having.

- How would you write a `while` statement in Go?

Go doesn't have a while statement but we can simulate one with a for loop:

```golang
package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
```

- What does the keyword `defer` do?

`defer` statement hold over the execution of the function until the deferred function returns.

- Does Go support pointers? How do arguments get passed around(by value or by reference)?

Yes. A pointer will hold the memory address of a value so arguments are passed by reference.

- Are arrays in Go fixed length? How about slices?

Arrays have a fixed size and all values share the same type, while slices are dynamically allocated

- Say you have a map: `map[string]int`, how would do a lookup and check to see if the map holds the value of the key you were looking for?
- How does Go structure programs? What is the difference between a library and a program that executes?
- How do make a function or a type public? And how do you make it private?
- You are going to be building a simple calculator with 4 basic operations(add, subtract, multiply and divide). First build a library that provides those 4 methods. After that implement a program that reads from the command line the operation to be done and prints the result(by calling the library you implemented previously). The operation should be read however you'd like, but for simplicity sake limit yourself to 2 operands and 1 operation character. Something like `./program 1 + 2`.
- Suppose you are building a web server that needs a DB that can do a set of simple operations. You know that the requirements of what DB to use will change. You also now that it will be easier for testing purposes to not have to setup something like MySQL. How would you solve this problem using the feature that Go provide?
- How would you build a simple function that can receive _any_ type of argument and prints the if that argument is of a primitive type. Limit to just `int`, `string`, `float` and `bool`.
- How are errors defined in Go?
- Ok, you know how errors are defined in Go now. Time to build a simple `errors` package that allows you to build errors that specify what kind of error is it, limit yourself to 3 kinds: `Internal`, `ThirdParty` and `Other`. Then provide a function in that package for users to check if the error they have is of the kind they care about. **NOTE**: remember not to break with the way errors are defined in Go, take advantage of that.
- What do you use to make two functions concurrent?
- How would you synchronize two concurrent functions?
- Write a program with three functions. One will send stuff(whatever you'd like) over a channel every one second and one will receive it and print it. The third function will tell the other two functions to stop and return(it could be the main func) after 5 seconds. **NOTE**: the program can not end until the sender and receiver have returned.
