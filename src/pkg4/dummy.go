// This package is only used for its 'init' function
package dummy

import display "pkg2"

func init() {
	display.ConsoleLogString("Only init is called for Dummy package")
}

// This will not be made available to 'main' package
func NotUsed() {
	display.ConsoleLogString("I am of no use")
}
