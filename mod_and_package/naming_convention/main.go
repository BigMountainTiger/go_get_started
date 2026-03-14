package main

import (
	"fmt"

	// the package name is different from the folder name
	// It is allowed but not recommended, the strong convention is to have the package name the same as the folder name
	// The import path is based on "the module name" / "the folder structure"
	any_package "song.com/go_get_started/mod_and_package/naming_convention/folder1"
)

func main() {
	fmt.Println("Hello, World!")

	// the package name is different from the folder name,
	// but we can still call the function using the package name
	any_package.Which_package()
}
