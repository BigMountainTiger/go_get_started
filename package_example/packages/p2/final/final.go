package final

import (
	"fmt"
)

func Final_method() {
	fmt.Println("From p2")
}

func GetString() string {
	return "It is from the p2/final package"
}

func init() {
	fmt.Println("init() in p2/final/final.go file")
	fmt.Println("init() in p2/final/final.go file is called only once, even though the package is imported from different packages")
	fmt.Println()
}
