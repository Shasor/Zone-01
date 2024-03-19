package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arglen := len(os.Args)
	for i := 1; i < arglen; i++ {
		str := os.Args[i]
		for _, letter := range str {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
