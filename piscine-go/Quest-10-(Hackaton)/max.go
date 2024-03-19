package piscine

func Max(a []int) int {
	var result int
	for i := 0; i < len(a); i++ {
		if a[i] > result {
			result = a[i]
		}
	}
	return result
}
