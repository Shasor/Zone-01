package piscine

func BasicAtoi(s string) int {
	runes := []byte(s)
	nbr := 0
	for i := range runes {
		nbr = (nbr * 10) + int(runes[i]-48)
	}
	return nbr
}
