package bdfs

import "sort"

/*
	solution 1 backtrace
		https://leetcode.com/problems/combination-sum/discuss/16502/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning)
*/

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	curNum := make([]int, 0)
	sort.Ints(candidates)
	backtrace(candidates, &curNum, &result, target, 0)
	return result
}

func backtrace(candid []int, curNum *[]int, result *[][]int, remain, start int) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		tmp := make([]int, len(*curNum))
		copy(tmp, *curNum)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(candid); i++ {
		*curNum = append(*curNum, candid[i])
		backtrace(candid, curNum, result, remain-candid[i], i)
		*curNum = (*curNum)[:len(*curNum)-1]
	}
}
