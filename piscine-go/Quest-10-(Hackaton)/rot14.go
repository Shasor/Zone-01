package piscine

func Rot14(s string) string {
	var result string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			newChar := char - 12
			if newChar < 'A' {
				newChar = newChar + 26
			}
			result += string(newChar)
		} else if char >= 'a' && char <= 'z' {
			newChar := char - 12
			if newChar < 'a' {
				newChar = newChar + 26
			}
			result += string(newChar)
		} else {
			newChar := char
			result += string(newChar)
		}
	}
	return result
}
