package main

import "fmt"

func main() {
	x := 1

	//'y' refer to x address on memory
	y := &x

	fmt.Println("Value of x is :", x)
	fmt.Println("Address of x is :", &x)
	fmt.Println("Value of y is :", y)

	*y = 100

	fmt.Println("Value of x is :", x)
	fmt.Println("Address of x is :", &x)
	fmt.Println("Value of y is :", y)

}
