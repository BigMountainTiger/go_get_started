package main

import (
	"fmt"
)

type I interface {
	print()
}

type T1 struct{}

func (v T1) print() {
	fmt.Println("Called through T1")
}

type T2 struct{}

func (v T2) print() {
	fmt.Println("Called through T2")
}

func polymorphism_example() {
	list := make([]I, 3, 30)

	list[1] = T1{}
	list[2] = T2{}

	// This is how slice appended
	list = append(list, nil)

	for i, v := range list {

		fmt.Printf("Item No.%v\n", i)
		if v == nil {
			fmt.Printf("No entry, exit\n\n")
			continue
		}

		v.print()

		fmt.Println()
	}

	fmt.Println("Polymorphism works in golang")
}
