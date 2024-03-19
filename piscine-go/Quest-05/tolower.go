package piscine

func ToLower(s string) string {
	sRune := []rune(s)
	for i, word := range sRune {
		if word >= 'A' && word <= 'Z' {
			sRune[i] = word + 32
		}
	}
	return string(sRune)
}
