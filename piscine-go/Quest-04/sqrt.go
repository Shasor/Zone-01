package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return nb
	}
	i := 2
	for ; i*i < nb; i++ {
	}
	if i*i == nb {
		return i
	}
	return 0
}
