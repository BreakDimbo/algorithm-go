package string

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: 暴力遍历
	solution 2: Sliding Window
	solution 3:

*/

/* sulution 1
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 0
	for i := range s {
		visited := make(map[byte]bool)
		count := traverse(i, s, visited, 0)
		if count > max {
			max = count
		}
	}
	return max
}

func traverse(i int, s string, visited map[byte]bool, count int) int {
	if i > len(s)-1 || visited[s[i]] {
		return count
	}
	count++
	visited[s[i]] = true
	return traverse(i+1, s, visited, count)
}
*/

/* solution 2
func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	i, j, ans, n := 0, 0, 0, len(s)

	for i < n && j < n {
		// extend the range [i,j]
		if !set[s[j]] {
			set[s[j]] = true
			j++
			ans = util.Max(ans, j-i)
		} else {
			delete(set, s[i])
			i++
		}
	}
	return ans
}
*/

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]int)
	n, i, j, ans := len(s), 0, 0, 0

	for i < n && j < n {
		// extend the range [i,j]
		if _, ok := set[s[j]]; ok {
			i = util.Max(i, set[s[j]])
		}
		ans = util.Max(ans, j-i+1)
		set[s[j]] = j + 1
		j++
	}
	return ans
}
