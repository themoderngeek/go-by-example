package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[0] = "zero"
	m[1] = "one"
	m[2] = "two"

	fmt.Println("Map:", m)

	one := m[1]

	fmt.Println("1 =", one)
	fmt.Println("map size:", len(m))

	delete(m, 2)
	fmt.Println("Delete 2:", m)

	_, prs := m[2]
	fmt.Println("Contains 2", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
