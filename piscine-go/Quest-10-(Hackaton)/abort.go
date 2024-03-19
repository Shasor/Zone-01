package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	return Sort(slice)[2]
}

func Sort(tab []int) []int {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab); j++ {
			if tab[i] < tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
	return tab
}
