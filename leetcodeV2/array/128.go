package array

import (
	"github.com/breakD/algorithms/leetcodeV2/util"
)

/*
	solution 1: sort
	solution 2:
		use hash
		check n-1 and n+1
		update boundary
*/

func longestConsecutive(nums []int) int {
	res := 0
	boundMap := make(map[int]int)
	for _, n := range nums {
		if _, ok := boundMap[n]; !ok {
			var left, right int
			if value, ok := boundMap[n-1]; ok {
				left = value
			}
			if value, ok := boundMap[n+1]; ok {
				right = value
			}
			sum := left + right + 1
			boundMap[n] = sum

			res = util.Max(res, sum)

			boundMap[n-left] = sum
			boundMap[n+right] = sum
		}
	}
	return res
}
