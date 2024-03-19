package piscine

func BasicAtoi2(s string) int {
	runes := []byte(s)
	nbr := 0
	for i := range runes {
		if runes[i] >= 48 && runes[i] <= 57 {
			nbr = (nbr * 10) + int(runes[i]-48)
		} else {
			return 0
		}
	}
	return nbr
}
