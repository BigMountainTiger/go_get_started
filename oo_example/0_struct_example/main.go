package main

import "fmt"

type child_struct struct {
	text string
}

type parent_struct struct {
	v_child child_struct
	p_child *child_struct
}

func main() {

	fmt.Println("0. A struct can be nested")

	v := parent_struct{
		v_child: child_struct{
			text: "Initial value of v_child",
		},
		p_child: &child_struct{
			text: "Initial value of p_child",
		},
	}

	fmt.Println()
	fmt.Println("1. Struct can be initiated at declaration")
	fmt.Println(v.v_child.text)
	fmt.Println(v.p_child.text)

	p := &parent_struct{
		v_child: child_struct{
			text: "Initial value of v_child",
		},
		p_child: &child_struct{
			text: "Initial value of p_child",
		},
	}

	fmt.Println()
	fmt.Println("2. A pointer to a struct can be initiated at declaration")
	fmt.Println(p.v_child.text)
	fmt.Println(p.p_child.text)
}
