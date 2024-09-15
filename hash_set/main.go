package main

import (
	"fmt"
	"go-dsa/hash_set/v2/hash_set"
)

type Person struct {
	Name string
	Age  int
}

// Hash returns the has of a person, it conforms to the hash_set.Hasher interface
func (p Person) Hash() string {
	return fmt.Sprintf("%s-%d", p.Name, p.Age)
}

func main() {
	mySet := hash_set.NewSet[Person, string]()

	// Add
	mySet.Add(Person{
		Name: "Batman",
		Age:  28,
	})
	mySet.Add(Person{
		Name: "Robin",
		Age:  16,
	})
	mySet.Add(Person{
		Name: "Batman",
		Age:  28,
	})

	// Print
	fmt.Printf("%s\n", mySet.String())

	// Contains
	result := mySet.Contains(Person{
		Name: "Batman",
		Age:  28,
	})
	fmt.Printf("Set contains batman %v\n", result)

	// Deletion
	mySet.Delete(Person{
		Name: "Batman",
		Age:  28,
	})
	fmt.Printf("%s\n", mySet.String())
}
