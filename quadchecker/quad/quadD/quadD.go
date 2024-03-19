package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		println("Vous devez fournir seulement deux arguments : X et Y")
		return
	}
	x, errX := strconv.Atoi(os.Args[1])
	y, errY := strconv.Atoi(os.Args[2])
	if errX != nil || errY != nil {
		println("Les arguments X et Y doivent être des nombres entiers")
		return
	}
	if x > 0 && y > 0 {
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		if x >= 2 {
			z01.PrintRune('C')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('\n')
		}
		for j := 0; j < y-2; j++ {
			z01.PrintRune('B')
			for k := 0; k < x-2; k++ {
				z01.PrintRune(' ')
			}
			if x >= 2 {
				z01.PrintRune('B')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('\n')
			}
		}
		if y > 1 {
			z01.PrintRune('A')
			for i := 0; i < x-2; i++ {
				z01.PrintRune('B')
			}
			if x >= 2 {
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('\n')
			}
		}
	} else {
		println("Les valeurs de x et y doivent être supérieures à zéro")
		return
	}
}
