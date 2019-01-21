package string

/*
	solution 1: brute force
	solution 2: KMP
*/

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		j, k := 0, i
		for ; j < len(needle); j++ {
			if k < len(haystack) && haystack[k] == needle[j] {
				k++
				continue
			} else {
				break
			}
		}
		if j == len(needle) {
			return k - j
		}

	}
	return -1
}
