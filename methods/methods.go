package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) parim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Rectangle:", r)

	fmt.Println("area:", r.area())
	fmt.Println("parim:", r.parim())

	rp := &r

	fmt.Println("area:", rp.area())
	fmt.Println("parim:", rp.parim())
}
