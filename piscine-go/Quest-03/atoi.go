package piscine

func Atoi(s string) int {
	isnegative := false
	runes := []byte(s)
	nbr := 0
	if len(s) == 0 {
		return 0
	}
	if runes[0] == 45 {
		isnegative = true
		runes = append(runes[:0], runes[1:]...)
	} else if runes[0] == 43 {
		runes = append(runes[:0], runes[1:]...)
	}
	for i := range runes {
		if runes[i] < 48 || runes[i] > 57 {
			return 0
		}
		nbr = (nbr * 10) + int(runes[i]-48)
	}
	if isnegative {
		return nbr * -1
	}
	return nbr
}
