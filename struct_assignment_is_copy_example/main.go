package main

import "fmt"

func assign_is_copy() {

	type s_example struct {
		value string
	}

	a := s_example{
		value: "Initial value",
	}

	b := a
	b.value = "New value"

	fmt.Println("a.value =", a.value)
	fmt.Println("b.value =", b.value)
	fmt.Println("When a struct assigned to another variable, a new instance is created")

	pa := &a
	pb := &b

	fmt.Println("The same pointer =", (pa == pb))
}

func string_assignment_is_copy() {
	a := "the string"
	b := a

	fmt.Println("string_assignment_is_copy", "The same pointer =>", &a == &b)
	fmt.Println("a different string is created during the copy,", "string is immutable")
}

func struct_assign_depends_if_pointer() {
	type s_internal struct {
		value string
	}

	i := s_internal{
		value: "Initial value",
	}

	type s_example struct {
		value *s_internal
	}

	a_obj := s_example{
		value: &i,
	}

	b_obj := a_obj
	b_obj.value.value = "Altered value"

	fmt.Println(a_obj.value.value, "-", b_obj.value.value)
	fmt.Println("Point to the same address =>", a_obj.value == b_obj.value)
}

func main() {
	line := "------------------------------"

	assign_is_copy()

	fmt.Println(line)
	string_assignment_is_copy()

	fmt.Println(line)
	struct_assign_depends_if_pointer()
}
