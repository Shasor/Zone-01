package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, strs := range tab {
		if f(strs) {
			x++
		}
	}
	return x
}
