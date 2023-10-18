package main

import "fmt"

type s_example struct {
	value string
}

func main() {

	a := s_example{
		value: "Initial value",
	}

	b := a
	b.value = "New value"

	fmt.Println("a.value =", a.value)
	fmt.Println("b.value =", b.value)
	fmt.Println("When a struct assigned to another variable, a new instance is created")

	pa := &a
	pb := &b

	fmt.Println("The same pointer =", (pa == pb))

}
