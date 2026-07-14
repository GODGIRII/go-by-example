package main

import (
	"fmt"
)

/*
	variables are declared explicitly
	var is used to declare 1 or more variables
	variables declared without a corresponding initializations are zero-valued

	Zero-values are the default value of any type
	Ex: Zero-value of int is 0

	:= is shorthand for declaring and initializing a variable
*/

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
