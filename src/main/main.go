package main

import (
	"fmt"
	// Using alias because directory name is different from the package name
	math "pkg1"
	// Nested package i.e. 'advanced' package lies inside 'math' package
	advanced "pkg1/pkg3"
	display "pkg2"

	// Blank identifier in import
	_ "pkg4"

	// Direct dependency package downloaded and installed using 'go get'
	"github.com/pborman/uuid"
)

// 'Init' function, which executes before 'main' function to initialize global variables
// Here, we are just printing a random string for demo purposes
/* 'Init' function is only called once for each package, irrespective of
how many times it gets imported */
func init() {
	display.ConsoleLogString("Hello from Package Main's init function")
}

func main() {
	display.ConsoleLogString("Starting program...")
	a := math.Add(1, 2)
	b := math.Subtract(1, 2)
	c := advanced.Multiply(1, 2)
	display.ConsoleLogInt(a)
	display.ConsoleLogInt(b)
	display.ConsoleLogInt(c)
	fmt.Print("UUID is: ")
	display.ConsoleLogString(uuid.New())
}
