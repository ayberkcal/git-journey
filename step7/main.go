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

func reverse(curr *Student) *Student {
	if curr.next == nil {
		return curr
	} else {
		newHead := reverse(curr.next)
		curr.next.next = curr
		curr.next = nil
		return newHead
	}
}

func main() {
	var students *Student
	students = nil

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
		newStudent := &Student{name: lineValues[0], surname: lineValues[1], age: tage, ssn: lineValues[3], next: students}
		students = newStudent
	}

	fmt.Println("\nStudent List:")
	for s := students; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}

	students = reverse(students)
	fmt.Println("\n Reserved Student List:")
	for s := students; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}

}
