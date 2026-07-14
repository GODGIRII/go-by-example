package main

import (
	"fmt"		//Import fmt package
)

func main () {

	fmt.Println("go" + "lang")		// Joined two string's with +

	/*
		Integers and floats
		Basic sums
	*/

	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	/*
		Booleans and Operators
	*/

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
