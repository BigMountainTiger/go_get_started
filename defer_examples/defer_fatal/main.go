package main

import (
	"fmt"
	"log"
)

func main() {

	defer func() {
		fmt.Println()
		fmt.Println("This is not called if Fatal() called")
		r := recover()

		if r != nil {
			log.Printf("recovered from Panic - %v", r)
		}
	}()

	// https://pkg.go.dev/log#Fatal
	fmt.Println("Fatal is equivalent to Print() followed by a call to os.Exit(1)")
	fmt.Println("There is no way to recover from \"os.Exit(1)\"")

	// log.Panic("Panic can be recovered")

	fmt.Println()
	log.Println("Calling log.Fatal()")
	log.Fatal("Fatal here")
}
