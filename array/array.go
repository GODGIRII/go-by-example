package main

import (
	"fmt"
)

func main() {

	var a [5]int	// Declaring size of array 5
	fmt.Println("emp:", a)	// Arrays are zero-valued to 0

	a[4] = 100	// Input value at index 4
	fmt.Println("set:", a)	//Print whole array
	fmt.Println("get:", a[4])	//Print just single index value

	fmt.Println("len:", len(a))	//Length

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int	//Multi-dimensional array
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
