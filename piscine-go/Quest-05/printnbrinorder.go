package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	tabint := [10]int{}
	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	} else {
		for {
			if n == 0 {
				break
			}
			tabint[n%10]++
			n /= 10
		}
		for i, nbr := range tabint {
			if nbr != 0 {
				for j := 0; j < nbr; j++ {
					z01.PrintRune(rune(i + 48))
				}
			}
		}
	}
}
