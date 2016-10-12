package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1

	fmt.Println("Inital:", i)

	zeroval(i)

	fmt.Println("set value with i:", i)

	zeroptr(&i)

	fmt.Println("i set by pointer", i)

	fmt.Println("i pointer:", &i)
}
