package math

// This package uses the 'display' package to log to standard output
import (
	"fmt"
	display "pkg2"
)

// 'Init' function
func init() {
	fmt.Println("Hello from Package Math's Add function")
}

func Add(a, b int) int {
	display.ConsoleLogString("Adding given numbers")
	return a + b
}
