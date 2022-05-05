package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "Tanner Stufflebeam"
	course := "Getting started"

	fmt.Println(converter(author, course))
}

func converter(author, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}