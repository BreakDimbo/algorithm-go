package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: stack
	计算每一个 bar 的左最远大于它的index和右最远
*/

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxArea := 0
	lessFromLeft := make([]int, len(heights))
	lessFromRight := make([]int, len(heights))
	lessFromLeft[0] = -1
	lessFromRight[len(heights)-1] = len(heights)

	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}

	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(heights) && heights[p] > heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}

	for i := range heights {
		maxArea = util.Max(maxArea, heights[i]*(lessFromRight[i]-lessFromLeft[i]-1))
	}
	return maxArea
}
