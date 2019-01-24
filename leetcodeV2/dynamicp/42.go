package dynamicp

import (
	"github.com/breakD/algorithms/leetcodeV2/util"
)

/*
	solution 1: brute force
	solution 2: dp min(leftMax-rightMax) - height
	solution 3: two pointers
*/

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	ans := 0
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = util.Max(leftMax[i-1], height[i])
	}

	for j := n - 2; j >= 0; j-- {
		rightMax[j] = util.Max(height[j], rightMax[j+1])
	}

	for i := range height {
		ans += util.Min(rightMax[i], leftMax[i]) - height[i]
	}
	return ans
}
