package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
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
			z01.PrintRune('C')
			for i := 0; i < x-2; i++ {
				z01.PrintRune('B')
			}
			if x >= 2 {
				z01.PrintRune('A')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}
