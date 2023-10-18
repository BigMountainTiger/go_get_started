package main

import (
	"fmt"
)

type s_example struct {
	number int
	value  string
}

func (v *s_example) get_name() string {
	return v.value
}

type i_example_interface interface {
	get_name() string
}

func main() {

	// Array assignment is copy
	var arr [3]int
	arr_copy := arr
	arr_copy[0] = 100
	fmt.Printf("0. Array assignment is copy %v, so value is not changed when the copy changes\n", arr[0])

	// slice assignment is pointer in nature
	arr_slice := arr[:]
	arr_slice[0] = 100
	fmt.Printf("0. Array assignment to a slice is pointer nature %v, so value is changed when the slice changes\n", arr[0])

	fmt.Println()

	// Make create a slice
	intArr := make([]int, 3)
	s_lice := intArr[:]
	s_lice[0] = 100

	fmt.Printf("0. Slice assignment change the pointer value => pointer equal = %v\n", &intArr == &s_lice)
	fmt.Printf("It seems like a struct having a pointer member\n")

	s_copy := intArr
	s_copy[1] = 100

	for i, v := range intArr {
		fmt.Printf("%v - %v\n", i, v)
	}

	intArrP := new([3]int)
	for i, v := range intArrP {
		fmt.Printf("%v - %v\n", i, v)
	}

	fmt.Println()

	fmt.Println("0. Make retures a slice, no an array, slice is pointer in nature")

	fmt.Println("1. Make retures a type, new returns a pointer")

	// Map is also pointer in nature
	m := make(map[string]int)
	m_a := m
	m_a["A"] = 1
	fmt.Printf("2. You can only Make a map - %v\n", m_a["A"])
	fmt.Printf("2. Map assignment is pointer in nature - %v\n", m["A"])

	// Use make for channel
	ch := make(chan int, 1)
	ch <- 100
	fmt.Printf("3. You can only Make a channel - %v\n", <-ch)

	// Interface is not assignable by a struct of the function is receiver is a pointer
	// Need to assign it to a pointer
	s := s_example{}

	fmt.Printf("4. Default value of struct members = %d - %v(emptry string)\n", s.number, s.value)

	var iface i_example_interface
	func() {
		s.value = "A name"

		iface = &s
		iface_p := &iface

		fmt.Printf("5. It is possible to have a pointer to an interface - %v\n", (*iface_p).get_name())
		fmt.Printf("5. It is possible to assign struct to the interface if the receiver is a pointer - %v\n", iface.get_name())
	}()

	fmt.Println()

	a_func := func() string {
		return "Value from the func()"
	}

	b_func := &a_func
	fmt.Printf("6. Func assignment has pointer nature %v - %v\n", (&a_func == b_func), (*b_func)())

	func() {

		iface = &s
		iface_2 := iface

		v_1 := iface.(*s_example)
		v_2 := iface_2.(*s_example)

		fmt.Printf("7. Pointer interface assert into the same pointer value => %v\n", v_1 == v_2)
	}()

	func() {

		fmt.Println()
		str := `First row - This is the first line
			Second row - This is the second line`

		fmt.Println(str)

		str = "First row - This is the first line\n" +
			"Second row - This is the second line"

		fmt.Println(str)

		fmt.Printf("8. Backtick support multi-line string, but has difficulty with indentations\n")
	}()

	i := 0
	func() {
		i = 100

		// g_i is in the package scope
		g_i = 100
	}()
	fmt.Println()
	fmt.Printf("9. It is allowed to modify variable defined in outer scope => %d, %d\n", i, g_i)
	fmt.Println("g_i is in the package scope")
}
