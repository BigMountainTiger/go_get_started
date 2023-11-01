package main

import "fmt"

func type_assertions() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	fmt.Println()

	s, ok := i.(string)
	fmt.Println(s, ok)
	fmt.Println()

	fmt.Println("With ok check, no panic, we know if it success by the ok value")
	f, ok := i.(float64)
	fmt.Println(f, ok)
	fmt.Println()

	fmt.Println("Without ok check, panic")
	f = i.(float64)
	fmt.Println(f)

}
