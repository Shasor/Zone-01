package piscine

func ForEach(f func(int), a []int) {
	for _, ints := range a {
		f(ints)
	}
}
