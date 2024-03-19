package piscine

func Unmatch(a []int) int {
	slice := make(map[int]int)
	for _, i := range a {
		slice[i]++
	}
	for _, i := range a {
		if slice[i]%2 == 1 {
			return i
		}
	}
	return -1
}
