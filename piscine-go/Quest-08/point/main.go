package main

import "github.com/01-edu/z01"

func setPoint(x, y *int) {
	*x = 42
	*y = 21
}

func main() {
	var pointsX int
	var pointsY int
	setPoint(&pointsX, &pointsY)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune(pointsX/10) + 48)
	z01.PrintRune(rune(pointsX%10) + 48)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune(pointsY/10) + 48)
	z01.PrintRune(rune(pointsY%10) + 48)
	z01.PrintRune('\n')
}
