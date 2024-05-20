package main

import "fmt"

type human struct{}
type dog struct{}

func (a human) speak() {
	fmt.Println("A human speaks")
}

func (a dog) speak() {
	fmt.Println("A dog barks")
}

type animal interface {
	speak()
}

func main() {
	var animals [2]animal
	animals[0] = human{}
	animals[1] = dog{}

	// golang supports polymorphism
	fmt.Println("We can implement polymorphism by interface")
	for _, a := range animals {
		a.speak()
	}
}
