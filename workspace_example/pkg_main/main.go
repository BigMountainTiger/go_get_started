package main

import (
	"fmt"

	"song.com/pkg_library/utils"
)

func main() {
	s := "abcdef"
	fmt.Println("Calling Reverse function from another local package")
	fmt.Println(s, utils.Reverse(s))
}
