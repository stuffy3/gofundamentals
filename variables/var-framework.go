package main

import (
	"fmt"
	"reflect"
	"strconv"
	"os"
)


func main() {
	name := os.Getenv("USERNAME")
	course := "Go Fundamentals"
	module := "4"
	clip := 2
	// courseComplete := false

	fmt.Println("Name and course are set to", name, "and", course, ".")
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

	var ptr *string = &course // pointers store the variable address of other variables, the asterisk is what says 
	//its a pointer and when used gives the value of the pointer not the address
	//putting & in from will return the variables address 
	fmt.Println("point course variable at address," ,ptr, "which holds this value,", *ptr)

}