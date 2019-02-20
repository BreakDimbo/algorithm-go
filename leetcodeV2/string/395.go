package string

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: sliding window
*/

func longestSubstring(s string, k int) int {
	var count int
	for h := 1; h <= 26; h++ {
		count = util.Max(count, longestSubStringWithUniqueRune(s, k, h))
	}
	return count
}

func longestSubStringWithUniqueRune(s string, k, targetUniqueNum int) int {
	m := make([]int, 26)
	noLessThanKNum := 0
	uniqueNum := 0
	start, end := 0, 0
	count := 0

	for end < len(s) {
		if m[s[end]-'a'] == 0 {
			uniqueNum++
		}
		m[s[end]-'a']++
		if m[s[end]-'a'] == k {
			noLessThanKNum++
		}
		end++

		for uniqueNum > targetUniqueNum {
			if m[s[start]-'a'] == k {
				noLessThanKNum--
			}
			m[s[start]-'a']--
			if m[s[start]-'a'] == 0 {
				uniqueNum--
			}
			start++
		}

		if uniqueNum == targetUniqueNum && uniqueNum == noLessThanKNum {
			count = util.Max(count, end-start)
		}
	}
	return count
}
