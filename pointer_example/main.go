package main

import "fmt"

type example_struct struct {
	number int
}

func print_separator() {
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println()
}

func assignment_to_a_variable_replaces_content_of_the_memory_block() {
	fmt.Println("Assignment to a non-pointer variable replaces the momory block of the same address")
	fmt.Println()

	es := example_struct{
		number: 0,
	}

	fmt.Println("Initial value", es)

	ptr := &es

	es = example_struct{
		number: 1,
	}

	fmt.Println("Assigned value", es)
	fmt.Println("Value accessed by pointer", *ptr)

	comment := "Assignment is replacing the content at the memory address (can be memory block)"
	fmt.Println(comment)

}

func assignment_to_a_dereferenced_pointer_replaces_content_of_the_memory_block() {
	fmt.Println("Assignment to dereferenced pointer replaces the momory block of the same address")
	fmt.Println()

	es := example_struct{
		number: 0,
	}

	fmt.Println("Initial value", es)

	initial_prt := &es

	// Deferencing assignment
	func(esp *example_struct) {

		*esp = example_struct{
			number: 1,
		}

		fmt.Println("dereferencing assignment is allowed", *esp)

	}(&es)

	fmt.Println("After dereferencing assignment, the original struct changed", es)
	after_assignment_ptr := &es
	fmt.Println("After dereferencing assignment, pointer address remains the same =", initial_prt == after_assignment_ptr)
	// Dereferencing assignment is replacing the content of the memory
	fmt.Println("*** dereferencing assignment is replacing the content of the memory")

	// Assignment through pointer
	func(esp *example_struct) {

		esp.number = 2
		fmt.Println("Assignment through pointer", *esp)
	}(&es)

	fmt.Println("After assignment through pointer", es)
}

func assignment_struct_receiver_get_a_copy() {
	fmt.Println("Assignment a struct, the receiver get a independent copy")
	fmt.Println()

	es := example_struct{
		number: 0,
	}

	es_1 := es
	es_2 := es
	es_3 := es

	es_1.number = 1
	es_2.number = 2
	es_3.number = 3

	fmt.Println(es)
	fmt.Println(es_1)
	fmt.Println(es_2)
	fmt.Println(es_3)
}

func declare_a_struct_get_all_default_value() {
	fmt.Println("Declare a struct without initialization get all default value")
	fmt.Println()

	var a example_struct
	fmt.Println(a)
}

func struct_assignment_copy_is_deep_copy() {
	fmt.Println("Struct assignment copy is deep copy")
	fmt.Println()

	type child struct {
		value int64
	}

	type example_struct struct {
		value child
	}

	a := example_struct{
		value: child{
			value: 0,
		},
	}
	b := a

	fmt.Println("Initial value of a", a)
	fmt.Println("Initial value of b", b)

	b.value.value++
	fmt.Println("After increment b.value.value, only b is affected, so it is a deep copy")
	fmt.Println("value of a", a)
	fmt.Println("value of b", b)
}

func main() {
	assignment_to_a_variable_replaces_content_of_the_memory_block()

	print_separator()
	assignment_to_a_dereferenced_pointer_replaces_content_of_the_memory_block()

	print_separator()
	assignment_struct_receiver_get_a_copy()

	print_separator()
	declare_a_struct_get_all_default_value()

	print_separator()
	struct_assignment_copy_is_deep_copy()
}
