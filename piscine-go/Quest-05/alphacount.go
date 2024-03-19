package piscine

func AlphaCount(s string) int {
	i := 0
	for _, v := range s {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			i++
		}
	}
	return i
}
