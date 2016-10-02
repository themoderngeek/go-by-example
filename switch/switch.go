package main

import "fmt"
import "time"

func main() {
	i := 2

	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's the Morning")
	default:
		fmt.Println("It's the Afternoon")
	}

	s := "fizz"

	switch s {
	case "fizz":
		fmt.Println("buzz")
	default:
		fmt.Println("fizz")
	}
}