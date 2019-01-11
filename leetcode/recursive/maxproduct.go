package recursive

import "math"

var max int

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max = math.MinInt32
	for i := range nums {
		product(nums, i, 1)
	}
	return max
}

func product(nums []int, i, curPro int) {
	if i >= len(nums) {
		return
	}
	curPro *= nums[i]
	if curPro > max {
		max = curPro
	}
	product(nums, i+1, curPro)
}
