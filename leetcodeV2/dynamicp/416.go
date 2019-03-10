package recursive

/*
	solution 1: dp 背包问题
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum >>= 1

	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true
	for i := 1; i < len(nums)+1; i++ {
		dp[i][0] = true
	}
	for j := 1; j < sum+1; j++ {
		dp[0][j] = false
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}
