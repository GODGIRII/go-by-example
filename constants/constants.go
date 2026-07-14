package main

import (
	"fmt"
	"math"
)

const s string = "constant" 	//const declares a constant value

/*
	const statement can also appear inside a function body

	constant expression perform arithmetic with arbitrary precision

	A numeric cosntant has no type until it's given one. such as by an explicit conversion
*/

func main() {

	fmt.Println(s)

	const n = 5000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
