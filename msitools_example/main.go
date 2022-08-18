package main

import "fmt"

// lay_1
func lay_1() {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println(r)
			fmt.Println("recover() takes effect on panics in the functon called by the current function")
		}
	}()

	fmt.Println("lay_1")

	lay_2()
}

// lay_2
func lay_2() {
	fmt.Println("lay_2")

	lay_3()
}

// lay_3
func lay_3() {
	fmt.Println("lay_3")

	a := 1
	b := 0

	c := a / b
	fmt.Println(c)
}

func main() {

	fmt.Println("Welcome to this msitools example application")

	lay_1()
}
