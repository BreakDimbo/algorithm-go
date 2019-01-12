package dynamic

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	max := 0

	for i := range nums {
		tmpMax := 0
		for j := 0; j <= i-1; j++ {
			if nums[i] > nums[j] && dp[j] > tmpMax {
				tmpMax = dp[j]
			}
		}
		dp[i] = tmpMax + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
