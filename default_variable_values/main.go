package main

import (
	"fmt"
)

type st struct{}
type ft func(v string) string

var i int
var f float64
var s string
var b bool
var slice []int
var mp map[string]int
var ch chan string
var pst *st
var ift ft

func main() {

	fmt.Println()
	fmt.Printf("1. Integer default value = %d\n", i)
	fmt.Printf("2. Float default value = %f\n", f)
	fmt.Printf("3. String default value is empty => %v\n", s == "")
	fmt.Printf("4. Bool default value = %v\n", b)

	fmt.Println()
	fmt.Printf("5. Slice default value is nil = %v\n", slice == nil)
	fmt.Printf("6. Map default value is nil = %v\n", mp == nil)
	fmt.Printf("7. Channel default value is nil = %v\n", ch == nil)
	fmt.Printf("8. Pointer default value is nil = %v\n", pst == nil)
	fmt.Printf("9. function default value is nil = %v\n", ift == nil)

	fmt.Println()
	fmt.Println("Total 9 types")
}
