package hash

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		if _, ok := m[v]; !ok {
			return false
		}
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
