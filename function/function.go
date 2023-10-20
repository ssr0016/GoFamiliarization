package main

//Functions are central in Go. We’ll learn about functions with a few different examples.

import "fmt"

func plus(a int, b int) int { // Here's a function takes two ints and returns thier sum as an int.
	return a + b
}

func plusPlus(a, b, c int) int { // When you have  consecutive parameters of the same type, you may omit the type name for the like-type
	return a + b + c // you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
}

func main() {
	res := plus(1, 2) // call a function just as you'd expect, with name (args). args means arguments
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

// $ go run functions.go
// 1+2 = 3
// 1+2+3 = 6
