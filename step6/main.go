package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name    string
	surname string
	age     int
	ssn     string
	next    *Student
}

func main() {

	//pointer of first student
	students := new(Student)

	students.next = nil

	if len(os.Args) < 2 {
		fmt.Println("There is no file")
		return
	}

	filename := os.Args[1]

	data, fileReadErr := ioutil.ReadFile(filename)

	if fileReadErr != nil {
		fmt.Println("somethink went wrong!")
	}

	lines := strings.Split(string(data), "\n")

	//range iterates over elements in a variety of data structures.
	for _, line := range lines {
		if line == "" {
			continue
		}

		lineValues := strings.Split(line, " ")
		tage, _ := strconv.Atoi(lineValues[2])
		newStudent := &Student{name: lineValues[0], surname: lineValues[1], age: tage, ssn: lineValues[3], next: students.next}
		students.next = newStudent
	}

	fmt.Println("Student List:")

	for s := students.next; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}

}
