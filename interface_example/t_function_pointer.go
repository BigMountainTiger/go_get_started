package main

import "fmt"

type HelloFunc func(string)

func get_func() (*HelloFunc, HelloFunc) {

	state := "It is possible to return a pointer to a function"

	f := HelloFunc(func(s string) {
		fmt.Println(s, state)
	})

	// return both the pointer to the function and the function itself
	return &f, f
}

func function_pointer() {
	pf, f := get_func()

	fmt.Println(pf, &f, "The pointer number changed")
	(*pf)("Called through the function pointer")
}
