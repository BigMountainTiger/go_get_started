package main

import (
	"errors"
	"fmt"
)

func basics() {
	// https://go.dev/ref/spec#Defer_statements
	fmt.Println("The basics:")
	fmt.Println("A \"defer\" statement invokes a function whose execution is deferred to the moment the surrounding function returns:")
	fmt.Println("1. because the surrounding function executed a return statement")
	fmt.Println("2. because reached the end of its function body,")
	fmt.Println("3. because the corresponding goroutine is panicking")
	fmt.Println()
}

func regular_return_value_function() string {

	r := "Initial value"

	defer func() {
		r = "Modified in a defer function"

		fmt.Println("defer function is called")
		fmt.Println("But the new value set in the defer function is not returned")
	}()

	return r
}

func named_return_value_function() (r string, e error) {

	fmt.Println("Mixing named and unnamed return parameter is a syntax error")

	r = "Initial value"
	e = nil

	defer func() {
		fmt.Println("The return parameters set by the return")
		fmt.Println(r)
		fmt.Println(e)

		fmt.Println("defer function is called")
		fmt.Println("the new values of the named return values will be returned")
		r = "Modified in a defer function"
		e = errors.New("An artificial error")
	}()

	// The return has the opportunity to set the named return parameters
	return "OK", errors.New("Error at return time")
}

func main() {
	basics()

	fmt.Println("Test regular_return_value_function:")
	r := regular_return_value_function()
	fmt.Println(r)
	fmt.Println()

	fmt.Println("Test named_return_value_function:")
	r, e := named_return_value_function()
	fmt.Println(r)
	fmt.Println(e)
	fmt.Println()

	fmt.Println("Conclusion:")
	fmt.Println("1. To modify the return parameters in a defer function, we MUST use named return parameters")
	fmt.Println("2. Mixing named and unnamed return parameters is a syntax error")
	fmt.Println("3. Functions of named return parameters have the return parameters pre-created before the functon is called")
	fmt.Println("4. Functions of named return parameters can have an empty return which effectively returns the values of the pre-created parameters")
	fmt.Println("5. If actual values are returned, the pre-created parameters are set by the actual values, and eventually returned to the caller")
	fmt.Println("6. The value last set in the defered function is effectively returned to the caller")
	fmt.Println()
}
