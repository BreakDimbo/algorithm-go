package string

import "strings"

/*
	solution 1: Horizontal scanning
	solution 2: Divide and conquer
*/

/*
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	return dividConquer(strs, 0, len(strs)-1)
}

func dividConquer(strs []string, lo, hi int) string {
	if lo == hi {
		return strs[lo]
	}

	mid := (lo + hi) / 2
	leftPrefix := dividConquer(strs, lo, mid)
	rightPrefix := dividConquer(strs, mid+1, hi)
	return commonPrefix(leftPrefix, rightPrefix)
}

func commonPrefix(lpre, rpre string) string {
	for strings.Index(lpre, rpre) != 0 {
		rpre = rpre[:len(rpre)-1]
		if len(rpre) == 0 {
			return ""
		}
	}
	return rpre
}
