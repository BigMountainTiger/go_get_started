package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	x := v.X
	y := v.Y

	return math.Sqrt(x*x + y*y)
}

func (v *Vertex) Set(x float64, y float64) {
	v.X = x
	v.Y = y
}

type Int int32

func (v *Int) Abs() float64 {
	fv := float64(*v)
	return math.Sqrt(fv * fv)
}

func Modify(v *Vertex) {
	v.X = 100
	v.Y = 100
}

func struct_example() {

	fmt.Println("You can modify the struct by pointer receivers")
	v := Vertex{3, 4}
	v.Set(3, 3)
	fmt.Println(v)
	fmt.Println()

	fmt.Println("Use method on structs")
	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	x := Int(-3)
	fmt.Println(x.Abs())
	fmt.Println()

	fmt.Println("Passing pointer to function")
	v = Vertex{3, 4}
	Modify(&v)
	fmt.Println(v)

}
