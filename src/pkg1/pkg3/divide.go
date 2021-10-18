// This is a nested package inside the 'math' package
package advanced

import (
	"fmt"
	display "pkg2"
)

func init() {
	fmt.Println("Hello from Package Math's Divide function")
}

func Multiply(a, b int) int {
	display.ConsoleLogString("Multiplying given numbers")
	return a * b
}
