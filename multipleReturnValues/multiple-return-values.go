package main

// Go has built-in support for multiple return values. This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.

import "fmt"

func vals() (int, int) { // The (int, int) in this function signature shows that the
	// function returns 2 int.
	return 3, 7
}

func main() {
	a, b := vals() //Here we use the 2 differen return values from the call with multiple assignments
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // if you only want a subset of the returned values, use blank identifier _.
	fmt.Println(c)
}

// $ go run multiple-return-values.go
// 3
// 7
// 7
