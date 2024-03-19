package piscine

func Any(f func(string) bool, a []string) bool {
	for _, strs := range a {
		if f(strs) {
			return true
		}
	}
	return false
}
