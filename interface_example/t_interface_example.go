package main

import (
	"fmt"
	"math"
)

type IMyFloat float64

func (f IMyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type IVertex struct {
	X, Y float64
}

func (v *IVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type IAbser interface {
	Abs() float64
}

func interface_example() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered")
			fmt.Println(r)
		}
	}()

	var a IAbser

	f := IMyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &f
	fmt.Println(a.Abs())

	// If the method receiver is on a pointer,
	// only a pointer can be assigned to an interfacce
	// the following wont compile
	// a = v

	a = &v
	fmt.Println(a.Abs())

	fmt.Println()
	func() {
		fmt.Println("Empty interface can hold values of any type")

		var i interface{}

		i = 12
		fmt.Printf("%T, %v\n", i, i)

		i = "Hello"
		fmt.Printf("%T, %v\n", i, i)
	}()

	fmt.Println()
	func() {
		// The following will panic
		fmt.Println("If an interface is nil, a method can be called with the reveiver value being nil")
		var vp *Vertex

		fmt.Println("vp is -", vp)
		a = vp
		vp.Abs()
	}()

}
