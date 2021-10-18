package math

import (
	"fmt"
	display "pkg2"
)

func init() {
	fmt.Println("Hello from Package Math's Subtract function")
}

func Subtract(a, b int) int {
	display.ConsoleLogString("Subtracting given numbers")
	return a - b
}
