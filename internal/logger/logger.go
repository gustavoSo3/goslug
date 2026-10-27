package logger

import (
	"fmt"
)

func PrintMenu() {
	fmt.Println("======================================================================================")
	fmt.Println("|\tGoSlug")
	fmt.Println("|\tWelcome")
	fmt.Println("|\thttps://github.com/gustavoSo3/goslug")
	fmt.Println("|\tv0.1b")
	fmt.Println("======================================================================================")
}

func Log(a ...any) {
	fmt.Print("LOG\t")
	fmt.Println(a...)
}

func Error(a ...any) {
	fmt.Print("ERROR\t")
	fmt.Println(a...)
}
