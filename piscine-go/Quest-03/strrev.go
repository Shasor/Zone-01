package piscine

func StrRev(s string) string {
	var b string
	for _, i := range s {
		b = string(i) + b
	}
	return b
}
