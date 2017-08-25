package main

import (
	"fmt"
)

type Student struct {
	name    string
	surname string
	Next    *Student
}

type Teacher struct {
	name    string
	surname string
	Next    *Student
}

func main() {
	ali := Student{"Ali", "Cevik", nil}

	veli := Student{"Veli", "Yılmaz", &ali}

	deniz := Teacher{"Deniz", "Turker", &veli}

	fmt.Println("The teacher is ", deniz.name)

	fmt.Printf("'%s' is student of '%s'", deniz.Next.name, deniz.name)
}
