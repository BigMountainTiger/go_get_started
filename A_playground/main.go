package main

import "fmt"

type st struct {
	text string
}

func main() {

	v := "Text 1"
	vp := &v

	v = "Text 2"

	fmt.Println(*vp)

}
