package string

/*
	solution 1: brute force
	solution 2: use hash
	solution 3: KMP
	solution 4: BM
*/

/*
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
*/

/*
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	m := make(map[string]int)

	for i := range haystack {
		if i+len(needle)-1 < len(haystack) {
			if _, ok := m[needle]; !ok {
				m[haystack[i:i+len(needle)]] = i
			}
		}
	}
	if i, ok := m[needle]; ok {
		return i
	}
	return -1
}
*/

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := generateNext(needle)
	j := 0

	for i := range haystack {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] + 1
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func generateNext(p string) []int {
	next := make([]int, len(p))
	next[0] = -1
	k := -1

	// 注意 i 从 1 开始
	for i := 1; i < len(p); i++ {

		// 找次长匹配前缀
		for k != -1 && p[i] != p[k+1] {
			k = next[k]
		}

		if p[i] == p[k+1] {
			k++
		}
		next[i] = k
	}
	return next
}
