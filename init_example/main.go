package main

import (
	"fmt"

	"song.com/go_get_started/init_example/a_package"
)

func main() {
	fmt.Println("func main() is called")
	a_package.Print()

	fmt.Println()
	fmt.Println("Conclusion")
	fmt.Println("1. In each package, the constants, variables, are initiated before the init() is called")
	fmt.Println("2. The init() is called before imported")
	fmt.Println("3. In each package, if more than one init(), they are called in lexical order by the file name")
}

func init() {
	fmt.Println("init() in main/main.go")
}
