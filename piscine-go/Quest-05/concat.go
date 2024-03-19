package piscine

func Concat(str1 string, str2 string) string {
	str1len := len(str1)
	return str1[:str1len] + str2
}
