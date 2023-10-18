package p2

import (
	"fmt"

	"song.com/go_get_started/package_example/packages/p2/final"
)

func From_p2() {
	s := final.GetString()
	fmt.Println("From p2 directly - " + s)
}

func GetString() string {
	return "It is from the p2 package"
}
