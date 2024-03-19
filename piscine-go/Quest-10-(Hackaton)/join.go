package piscine

func Join(strs []string, sep string) string {
	result := ""
	if len(strs) == 0 {
		return "\n"
	}
	for i := 0; i < len(strs); i++ {
		result = result + strs[i] + sep
	}
	result = result[:len(result)-len(sep)]
	return result
}
