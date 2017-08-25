package main

import (
	"fmt"
)

func main() {
	x := fib(5)
	fmt.Printf("Fib is: '%d'", x)
}

func fib(x int) int {
	if x <= 3 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}
