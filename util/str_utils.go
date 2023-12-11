package util

func Concat(strs ...string) string {
	var result string
	for _, str := range strs {
		result += str
	}
	return result
}
