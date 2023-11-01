package main

import "fmt"

type typeA struct {
	text string
}

// Value receiver method
func (t typeA) value_receiver_method(pt *typeA) {
	fmt.Printf("With value receiver, same object = %v\n", &t == pt)
}

// Pointer receiver method
func (t *typeA) pointer_receiver_method(pt *typeA) {
	fmt.Printf("With pointer receiver, same object = %v\n", t == pt)
}

func main() {

	fmt.Println()
	t := typeA{text: "Initial text"}
	pt := &t

	fmt.Println("Experiment:")
	t.value_receiver_method(&t)
	pt.value_receiver_method(pt)

	t.pointer_receiver_method(&t)
	pt.pointer_receiver_method(pt)

	fmt.Println()
	fmt.Println("Conclusion:")
	fmt.Println("1. A receiver method can be called through either a value of pointer type, regardless if it is declared on a value or a pointer")
	fmt.Println("2, If the method is defined on a value, a copy of the calling value or a copy of the value pointed by the pointer is passed")
	fmt.Println("3. If the method is defined on a pointer, the pointer of the same object is passed to the receiving variable")
	fmt.Println()

	fmt.Println("Additional note:")
	fmt.Println("https://go.dev/tour/methods/8")
	fmt.Println("1. If there is an intention to modify the recevier value, we need to use a pointer receiver")
	fmt.Println("2. A pointer receiver does not involve coying object, it can be more efficient")
	fmt.Println("3. It is recommended (by golang document) not to mix value and pointer receivers on the same type")
	fmt.Println()
}
