package main

import "fmt"

func message() (string, string) {
	return "hello", "world"
}

func main() {
	a, b := message()
	fmt.Println(a, b)

	c, _ := message()
	fmt.Println(c)

	_, d := message()
	fmt.Println(d)
}
