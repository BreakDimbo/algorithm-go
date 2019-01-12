package dynamic

func coinChange(coins []int, amount int) int {
	max := amount + 1

	// dp[i] represent the min coins for target value i
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = max
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			// 注意：需要当前目标值大于硬币面值
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	// not found solution
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
