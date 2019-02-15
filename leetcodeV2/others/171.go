package others

/*
	solution 1: 26进制
*/

func titleToNumber(s string) int {
	var res int
	for _, c := range s {
		res = res*26 + (int(c-'A') + 1)
	}
	return res
}
