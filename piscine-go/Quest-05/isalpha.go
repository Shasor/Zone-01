package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) || s[i] == 35 || s[i] == 42 || s[i] == 64) {
			return false
		}
	}
	return true
}
