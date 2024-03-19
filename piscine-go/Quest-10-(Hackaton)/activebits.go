package piscine

func ActiveBits(n int) int {
	i := 0
	for n != 0 {
		if n%2 == 1 {
			i++
		}
		n /= 2
	}
	return i
}
