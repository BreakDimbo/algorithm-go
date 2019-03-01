package string

/*
	solution 1 : sliding window
	https://leetcode.com/problems/find-all-anagrams-in-a-string/discuss/92007/Sliding-Window-algorithm-template-to-solve-all-the-Leetcode-substring-search-problem.
*/

func findAnagrams(s string, p string) []int {
	m := make(map[byte]int)
	for i := range p {
		m[p[i]]++
	}
	result := make([]int, 0)
	count := len(m)

	begin, end := 0, 0

	for end < len(s) {
		if _, ok := m[s[end]]; ok {
			m[s[end]]--
			if m[s[end]] == 0 {
				count--
			}
		}
		end++

		for count == 0 {
			if _, ok := m[s[begin]]; ok {
				m[s[begin]]++
				if m[s[begin]] > 0 {
					count++
				}
			}
			if end-begin == len(p) {
				result = append(result, begin)
			}
			begin++
		}
	}
	return result
}
