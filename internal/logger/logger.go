package logger

import (
	"fmt"
)

// Prints GoSlug Menu
func PrintMenu() {
	fmt.Println("======================================================================================")
	fmt.Println("   @_@       ──► GoSlug")
	fmt.Println("  (   )====> ──► Welcome")
	fmt.Println("  `---'      ──► https://github.com/gustavoSo3/goslug | v0.1b")
	fmt.Println("======================================================================================")
}

// Used to log General data to the console
func Log(a ...any) {
	fmt.Print("LOG\t")
	fmt.Println(a...)
}

// Used to log Error data to the console
func Error(a ...any) {
	fmt.Print("ERROR\t")
	fmt.Println(a...)
}
