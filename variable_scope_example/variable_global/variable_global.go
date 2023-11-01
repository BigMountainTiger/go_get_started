package variable_global

import "fmt"

var g_variable string = "Initial - A global string"
var G_variable string = "Capitalized variable"

func Global_variable_is_global() {

	fmt.Println(g_variable)
	g_variable = "Updated - A global string"
	fmt.Println(g_variable)

	fmt.Println("Global variables:")
	fmt.Println("1. Global variables are declare out of a function")
	fmt.Println("2. A global variable exists throughout the lifetime of a program")
	fmt.Println()
}

func Global_variable_is_hidden_by_local_variable() {

	g_variable := "Locally declare value by :="
	fmt.Printf("The locally declared variable is used - %v", g_variable)
	fmt.Println()
}

func Global_variable_Print_the_global() {
	fmt.Printf("The global variable is not changed - %v", g_variable)
	fmt.Println()
}
