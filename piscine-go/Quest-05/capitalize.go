package piscine

func Capitalize(s string) string {
	rune_s := []rune(s)
	slen := len(s)
	if rune_s[0] >= 'a' && rune_s[0] <= 'z' {
		rune_s[0] = rune(rune_s[0] - 32)
	}
	for i := 1; i < slen; i++ {
		if (rune_s[i-1] < 'A' || rune_s[i-1] > 'Z') && (rune_s[i] >= 'a' && rune_s[i] <= 'z') && (rune_s[i-1] < 'a' || rune_s[i-1] > 'z') && (rune_s[i-1] < '0' || rune_s[i-1] > '9') {
			rune_s[i] = rune(rune_s[i] - 32)
		}
		if (rune_s[i-1] >= 'A' && rune_s[i-1] <= 'Z') && (rune_s[i] >= 'A' && rune_s[i] <= 'Z') {
			rune_s[i] = rune(rune_s[i] + 32)
		}
		if (rune_s[i-1] >= 'a' && rune_s[i-1] <= 'z') && (rune_s[i] >= 'A' && rune_s[i] <= 'Z') {
			rune_s[i] = rune(rune_s[i] + 32)
		}
		if (rune_s[i-1] >= '0' && rune_s[i-1] <= '9') && (rune_s[i] >= 'A' && rune_s[i] <= 'Z') {
			rune_s[i] = rune(rune_s[i] + 32)
		}
	}
	return string(rune_s)
}
