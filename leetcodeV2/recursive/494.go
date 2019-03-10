package recursive

/*
	solution 1: recursive O(2^n)
	solution 2: dp
		sum(P) - sum(N) = target
		sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
		2 * sum(P) = target + sum(nums)
		leetcode 416
*/

func findTargetSumWays(nums []int, S int) int {
	count := 0
	helper(nums, S, 0, &count)
	return count
}

func helper(nums []int, target, cur int, count *int) {
	if len(nums) == 0 {
		if target == cur {
			*count++
		}
		return
	}
	helper(nums[1:], target, cur+nums[0], count)
	helper(nums[1:], target, cur-nums[0], count)
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
