package pass

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
