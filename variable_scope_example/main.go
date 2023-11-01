package main

import (
	"fmt"

	"song.com/variable_scode_example/variable_global"
	"song.com/variable_scode_example/variable_local"
)

func main() {

	fmt.Println("Global Variable Test:")
	variable_global.Global_variable_is_global()
	variable_global.Global_variable_is_hidden_by_local_variable()
	variable_global.Global_variable_Print_the_global()
	fmt.Printf("Only capitalize named global variable accessible outside of the package - %v\n", variable_global.G_variable)
	fmt.Println()

	fmt.Println("Local Variable Test:")
	variable_local.Local_variable_is_block_structured()
}
