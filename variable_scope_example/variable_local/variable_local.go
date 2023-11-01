package variable_local

import (
	"fmt"
)

func Local_variable_is_block_structured() {
	fmt.Println("1. Local variables are declated inside a function")
	fmt.Println("2. Local variables in go are block scoped")
	fmt.Println("3. A pair of curly braces (may have short hand notations) defines a block")
	fmt.Println("4. Variables declared in outter scope is accessible in a inner block, but not vice versa")
	fmt.Println("5. Once out of scope a local variable is no longer accessible")
}
