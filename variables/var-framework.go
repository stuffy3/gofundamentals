package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name = "Tanner"
	course = "Go Fundamentals"
	module = "4"
	clip = 2
)
func main() {
	fmt.Println("Name and course are  set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("name is of type", reflect.TypeOf(name))
	fmt.Println("module is of type", reflect.TypeOf(module))
	// total := module + clip
	// fmt.Println("Module plus clip equals", total)

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}
}