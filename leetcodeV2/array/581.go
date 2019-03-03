package array

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: sorting
	solution 2: use stack
*/

/*
func findUnsortedSubarray(nums []int) int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	start, end := 0, 0
	for i, v := range sortedNums {
		if v != nums[i] {
			start = util.Min(start, i)
			end = util.Max(end, i)
		}
	}

	if end-start >= 1 {
		return end - start + 1
	}
	return 0
}
*/

func findUnsortedSubarray(nums []int) int {
	stack := make([]int, 0, len(nums))
	l, r := len(nums)-1, 0

	for i := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			l = util.Min(l, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			r = util.Max(r, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	if r-l >= 1 {
		return r - l + 1
	}
	return 0
}
