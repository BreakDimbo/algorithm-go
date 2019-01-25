package bdfs

/*
	solution 1: DFS
	solution 2: dp
	solution 3: greedy, rigth-to-left, check whether we can reach leftMostIndex
*/

/*
func canJump(nums []int) bool {
	return canJumpTo(nums, 0)
}

func canJumpTo(nums []int, i int) bool {
	if i == len(nums)-1 {
		return true
	}

	furthestJump := util.Min(i+nums[i], len(nums)-1)
	for j := furthestJump; j > i; j-- {
		if canJumpTo(nums, j) {
			return true
		}
	}

	return false
}
*/

/*
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[n-1] = true

	for i := n - 2; i >= 0; i-- {
		furthestJump := util.Min(nums[i]+i, n-1)
		for j := i + 1; j <= furthestJump; j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
*/

func canJump(nums []int) bool {
	n := len(nums)
	leftMostIndex := n - 1

	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= leftMostIndex {
			leftMostIndex = i
		}
	}
	return leftMostIndex == 0
}
