package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := (os.Args[0])[2:]
	for _, letter := range str {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
