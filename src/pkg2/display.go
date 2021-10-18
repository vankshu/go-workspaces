// Name of the package is not the directory name
// Package name is what we use in the package declaration
// Directory name is used to find the package location
package display

import "fmt"

// Exported functions as the name is capitalized
func ConsoleLogString(str string) {
	fmt.Println(str)
}

func ConsoleLogInt(num int) {
	fmt.Println(num)
}
