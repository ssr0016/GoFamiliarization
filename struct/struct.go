package main

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person { //newPerson constructs a new person struct with the given name.

	p := person{name: name} //You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20}) // This syntax creates a new struct.

	fmt.Println(person{name: "Alice", age: 30}) //You can name the fields when initializing a struct.

	fmt.Println(person{name: "Fred"}) //  Omitted fields will be zero-valued.

	fmt.Println(&person{name: "Ann", age: 40}) //An & prefix yields a pointer to the struct.

	fmt.Println(newPerson("jon")) // It’s idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50} //Access struct fields with a dot.
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) //You can also use dots with struct pointers - the pointers are automatically dereferenced.

	sp.age = 51
	fmt.Println(sp.age) //Structs are mutable.

	dog := struct { //If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}

// $ go run structs.go
// {Bob 20}
// {Alice 30}
// {Fred 0}
// &{Ann 40}
// &{Jon 42}
// Sean
// 50
// 51
// {Rex true}
