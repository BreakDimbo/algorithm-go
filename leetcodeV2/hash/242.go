package hash

/*
	solution 1: 使用 map
	solution 2: 使用数组：rune -> ASCII int(r - '0')
	solution 3: 排序，比较俩个字符串是否相等
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	asciiMap := make([]int, 256)
	for i := range s {
		asciiMap[int(s[i]-'0')]++
		asciiMap[int(t[i]-'0')]--
	}

	for _, v := range asciiMap {
		if v != 0 {
			return false
		}
	}
	return true
}
