// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int) // To create an empty map, use the built-in make: make(map[key-type]val-type)

	m["k1"] = 7 // set key/value pairs using typical name[key]=val syntax
	m["k2"] = 13

	fmt.Println("map:", m) // Printing a map with e.g fmt, Println will show all of its key/value pairs.

	v1 := m["k1"] // Get a value for a key with name[key]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) //if the key doens't exist, the zero value of the value type is returned

	fmt.Println("len:", len(m)) //the builtin len returns the number of key/value pairs when called on a map

	delete(m, "k2")
	fmt.Println("map:", m) //the builtin delete removes key/value pairs from a map.

	clear(m)
	fmt.Println("map:", m) // To remove all key/value pairs from a map, use the clear builtin

	_, prs := m["k2"]
	fmt.Println("prs:", prs) //The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate
	// between missing keys and keys with zero values like 0 or "". Here we didn’t need
	// the value itself, so we ignored it with the blank identifier _.

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // you can also declare and initialize a new map in the same line with this syntax

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
