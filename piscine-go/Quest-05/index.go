package piscine

func Index(s string, toFind string) int {
	slen := len(s)
	toFindlen := len(toFind)
	for i := 0; i < slen; i++ {
		if toFindlen != 0 && s[i] == toFind[0] {
			found := true
			ln3 := 0
			for j := 0; j < toFindlen; j++ {
				if i+ln3 >= slen || toFind[j] != s[i+ln3] {
					found = false
					break
				}
				ln3++
			}
			if found {
				return i
			}
		}
	}
	if toFind == "" {
		return 0
	}
	return -1
}
