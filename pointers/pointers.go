package main

// Go supports pointers, allowing you to pass references to values and records within your program.

import "fmt"

func zeroval(ival int) { //We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter,
	ival = 0 //so arguments will be passed to it by value. zeroval will get a copy of ival distinct from the one in the calling function.

}

func zeroptr(iptr *int) { //zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
	*iptr = 0 // The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
	// Assigning a value to a dereferenced pointer changes the value at the referenced address.
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval")

	zeroptr(&i)
	fmt.Println("zeroptr:", i) //The &i syntax gives the memory address of i, i.e. a pointer to i.

	fmt.Println("pointer:", &i) //Pointers can be printed too.
}

//$ go run pointers.go
// initial: 1
// zeroval: 1
// zeroptr: 0
// pointer: 0x42131100
