package main

import (
	"fmt"
	"math"
	"reflect"
)

// MyVFloat has a implementation on value
// MyPFloat has a implementation on pointer
type MyVFloat float64
type MyPFloat float64

func (f MyVFloat) Abs() float64 {
	v := float64(f)
	return math.Sqrt(v * v)
}

func (f *MyPFloat) Abs() float64 {
	v := float64(*f)
	return math.Sqrt(v * v)
}

type Abser interface {
	Abs() float64
}

func main() {

	fmt.Println()
	fmt.Println("The empty interface")
	fmt.Println("An empty interface may hold values of any type")
	var ei interface{}
	ei = "A string assigned to an empty interface"
	fmt.Printf("(%v, %T)\n", ei, ei)

	fmt.Println()
	fmt.Println("Non-empty interface")
	fmt.Println("1. An interface type is defined as a set of method signatures")
	fmt.Println("2. If the method is implemented on a value, both a pointer and a value can be assigned to the interface")
	fmt.Println("3. If the method is implemented on a pointer, only a pointer can be assigned to the interface")
	fmt.Println("4. It is recommended (by golang document) not to mix value and pointer receivers on the same type")

	// Assign to interface with value receiver method
	fmt.Println()
	fmt.Println("Example on value receiver")

	fv := MyVFloat(-1)
	var iv Abser = fv
	var ip Abser = &fv

	fmt.Printf("Calling from value = %v\n", iv.Abs())
	fmt.Printf("Calling from pointer = %v\n", ip.Abs())

	// Assign to interface with pointer receiver method
	fmt.Println()
	fmt.Println("Example on pinter receiver")

	fp := MyPFloat(-1)
	// iv = fp
	ip = &fp

	fmt.Printf("A value is not assignable to the interface if the method is implemented on a pointer\n")
	fmt.Printf("Calling from pointer = %v\n", ip.Abs())

	fmt.Println()
	fmt.Println("Type assertion")
	fmt.Println("https://go.dev/tour/methods/15")

	v, ok := ip.(*MyPFloat)
	fmt.Println("1. A type assertion provides access to an interface value's underlying concrete value")
	fmt.Println("2. The ok return value shows if the assertion successful")
	fmt.Println("3. If ok is not expected and if assertion fails, it is a panic")
	fmt.Printf("The type of v is %v, assertion is %v\n", reflect.TypeOf(v), ok)

}
