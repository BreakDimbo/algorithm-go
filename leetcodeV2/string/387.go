package string

/*
	solutiuon 1: brute force
	solution 2: hash
*/

func firstUniqChar(s string) int {
	m := make([]int, 26)

	for _, r := range s {
		m[r-'a']++
	}

	for i, r := range s {
		if m[r-'a'] == 1 {
			return i
		}
	}
	return -1
}
