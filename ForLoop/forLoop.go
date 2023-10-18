package main

// for is Go’s only looping construct. Here are some basic types of for loops.
//A classic initial/condition/after for loop.
// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
//You can also continue to the next iteration of the loop.

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
