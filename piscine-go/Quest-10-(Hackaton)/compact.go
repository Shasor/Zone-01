package piscine

func Compact(ptr *[]string) int {
	var result []string
	var i int
	for _, str := range *ptr {
		if str != "" {
			result = append(result, str)
			i++
		}
	}
	*ptr = result
	return i
}
