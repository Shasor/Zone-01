package piscine

func ConcatParams(args []string) string {
	result := args[0]
	for _, arg := range args[1:] {
		result += "\n" + arg
	}
	return result
}
