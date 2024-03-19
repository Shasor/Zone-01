package piscine

func BasicJoin(elems []string) string {
	var str string
	for _, e := range elems {
		str = str + e
	}
	return str
}
