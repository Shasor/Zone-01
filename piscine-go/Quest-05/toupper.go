package piscine

func ToUpper(s string) string {
	sRune := []rune(s)
	for i, word := range sRune {
		if word >= 'a' && word <= 'z' {
			sRune[i] = word - 32
		}
	}
	return string(sRune)
}
