package main

import (
	"fmt"

	p1 "song.com/go_get_started/package_example/packages/p1/final"
	"song.com/go_get_started/package_example/packages/p2"
	p2_final "song.com/go_get_started/package_example/packages/p2/final"
	"song.com/go_get_started/package_example/packages/p3"
)

func main() {
	fmt.Println("Start")

	p1.Final_method()
	p2_final.Final_method()
	p2.From_p2()
	p3.Final_method()

	fmt.Println()
	fmt.Println("0. go imports packages, and resources are used through the packages")
	fmt.Println("1. Package name is the fully qualified name, including the directory structure")
	fmt.Println("2. When the short name conflicting, alias can be used")

	fmt.Println()
	fmt.Println("3. Circular importing is not allowed in go")
	fmt.Println("4. The run-time will load/initialize a package only once even when it is imported multiple times")
}
