package main

import "fmt"

// Creating interface named Animal
// The idea of an interface is that it provides to us the methods that we need to implement
// They are named colletions of method signitures
type Animal interface {
	//Methods of the interface
	speak() string
	Legs() int
	Pet() bool
}

//Note that to implement an interface or for it to be considered an animal in this case, we need to implement ALL METHODS in the interface

// Creating animals as structs for our interface
type Dog struct {
	Breed string
	Age   int
	Color string
}

type Bear struct {
	Name  string
	Age   int
	Color string
}

type Lizard struct {
	Size  string
	Color string
	Name  string
}

// Like mentioned the type dog is considered a struct because it contains all methods of the interface
func (d Dog) speak() string {
	return "ruff"
}

func (d Dog) Legs() int {
	return 4
}

func (d Dog) Pet() bool {
	return true
}

// Same goes for bear, for it to be considered an animal it needs to have all the methods of the interface
func (b Bear) speak() string {
	return "roar"
}

func (b Bear) Legs() int {
	return 2
}

func (b Bear) Pet() bool {
	return false
}

// Same goes for lizard
func (l Lizard) speak() string {
	return "sss"
}

func (l Lizard) Legs() int {
	return 4
}

func (l Lizard) Pet() bool {
	return true
}

// Function to print the animal along with the interfaces methods
// We know we would pass in an Animal for the parameters in the main
func PrintAll(a Animal) {
	fmt.Println("This animal speaks ", a.speak(), " and has ", a.Legs(), "legs. This animal is", a.Pet(), " a pet")
}

// Specific function to print the bear name
func PrintBearName(b Bear) (string, string) {
	return "This bear name is: ", b.Name
}

func main() {

	//Creating the 'objects' of the animals so we could pass it to our print all function. In this case we create the dog object.
	animal1 := Dog{
		Breed: "Labrador",
		Age:   6,
		Color: "Brown",
	}
	//Creating object for the bear
	animal2 := Bear{
		Name:  "Carlos",
		Age:   14,
		Color: "Black",
	}
	//Creating the object for Lizard
	animal3 := Lizard{
		Size:  "big",
		Color: "Yellow",
		Name:  "Kevin",
	}

	//Calling PrintAll function to display and passing in the animals to them
	PrintAll(animal1)
	PrintAll(animal2)
	PrintAll(animal3)
	//Printing out just the bear name
	fmt.Println(PrintBearName(animal2))

}
