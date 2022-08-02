package main

import (
	"fmt"
	"log"
)

// The example of altering the return value in deferred function
func c() (s string) {

	defer func() { s = fmt.Sprintf("%s altered", s) }()

	return "Initial string"
}

func main() {

	fmt.Println(c())
	fmt.Println("Defered functon can alter the return value if needed")
	fmt.Println()

	log.Println("Start ...")

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic - detail: %s\n", r)
		}

		fmt.Println("Note:")
		fmt.Println("1. Deferred functions are called in LIFO order")
		fmt.Println("2. Deferred functions are called even when panic situations")
		fmt.Println("3. Panic can happen in deferred functions")
		fmt.Println("4. If panics happen in multiple deferred functions, the last one will be received by the recover() function")
		fmt.Println("5. If panics are recovered, the program continues as normals")

	}()

	defer func() {
		log.Println("Defer No.1")
		log.Panicln("Panic Here in No.1")
	}()

	defer func() {
		log.Println("Defer No.2")
		log.Panicln("Panic Here in No.2")
	}()

	log.Panicln("Panic in normal")
}
