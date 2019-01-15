package hash

/*
	solution 1: 使用 map
	solution 2: 使用数组：rune -> ASCII int(r - '0')
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	asciiMap := make([]int, 256)
	for _, v := range s {
		asciiMap[int(v-'0')]++
	}

	for _, v := range t {
		if asciiMap[int(v-'0')] < 1 {
			return false
		}
		asciiMap[int(v-'0')]--
	}

	for _, v := range asciiMap {
		if v != 0 {
			return false
		}
	}
	return true
}
