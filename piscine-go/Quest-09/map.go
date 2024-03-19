package piscine

func Map(f func(int) bool, a []int) []bool {
	var bools []bool
	for _, ints := range a {
		x := f(ints)
		bools = append(bools, x)
	}
	return bools
}
