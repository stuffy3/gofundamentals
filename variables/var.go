package main

import (
	"fmt"
)

func main() {
	name := "Tanner Stufflebeam"
	course := "Getting Started With Go"

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(course)
	fmt.Println("You're currently Watching", course)
}

func updateCourse(course string) string {
	course = "Getting Started with Docker"
	fmt.Println("Updated course to" course)
	return course
}