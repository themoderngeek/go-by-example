package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("Empty slice: ", s)

	s[0] = "hello"
	s[1] = "world"
	s[2] = "!"

	fmt.Println("Full slice: ", s)
	fmt.Println("element 1: ", s[1])

	s = append(s, "from")
	s = append(s, "go", "by", "example")

	fmt.Println("Appended slice: ", s)
	fmt.Println("Length of slice: ", len(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copy of slice s: ", c)

	l := s[3:5]
	fmt.Println("elements 3 - 5: ", l)

	l = s[:5]
	fmt.Println("all elements to 5: ", l)

	l = s[2:]
	fmt.Println("Elements from 2 onwards", l)

	f := []string{"d", "i", "l"}

	fmt.Println("Declared in line: ", f)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	naughtsAndCrosses := [][]string{}

	row1 := []string{"O", "X", "X"}
	row2 := []string{"X", "O", "O"}
	row3 := []string{"X", "O", "O"}

	naughtsAndCrosses = append(naughtsAndCrosses, row1)
	naughtsAndCrosses = append(naughtsAndCrosses, row2)
	naughtsAndCrosses = append(naughtsAndCrosses, row3)

	fmt.Println(naughtsAndCrosses[0])
	fmt.Println(naughtsAndCrosses[1])
	fmt.Println(naughtsAndCrosses[2])
}
