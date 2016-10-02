package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100

	fmt.Println("Setting a[4] to 100", a)
	fmt.Println("Get a[4]: ", a[4])

	fmt.Println("The length of a is", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Creating new array b", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2 dimensional arrray: ", twoD)
}
