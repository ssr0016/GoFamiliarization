package main

import "fmt"

// Go supports anonymous functions, which can form closures. Anonymous functions are useful
// when you want to define a function inline without having to name it.

func intSeq() func() int { // This functions intSeq return another function, which we define anonymously
	i := 0 //in the body of intSeq. The returned function closes over the variable i to form a closure.
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq() // We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be
	// upadated each time call nextInt.

	fmt.Println(nextInt()) // see the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // To confirm that state is unique to that particular function create and test a new one.
	fmt.Println(newInts())

}

// $ go run closures.go
// 1
// 2
// 3
// 1
