package piscine

func TrimAtoi(s string) int {
	rune_s := []rune(s)
	nbr := 0
	sign := 1
	c := 0
	minus_exists := false
	minus_pos := 0
	first_nbr := 0
	for i, char := range rune_s {
		if char == '-' {
			minus_pos = i
			minus_exists = true
		}
	}
	for _, word := range rune_s {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				c++
			}
			nbr = nbr*10 + c
			c = 0
		}
	}
	if minus_exists {
		for i, char := range rune_s {
			if char >= '0' && char <= '9' {
				first_nbr = i
				break
			}
		}
	}
	if first_nbr > minus_pos {
		sign = -1
	}
	return nbr * sign
}
