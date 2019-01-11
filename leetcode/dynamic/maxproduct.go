package dynamic

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i][min: 1]
	// dp[i][max: 0]
	dp := [][]int{
		{nums[0], nums[0]},
		{0, 0},
	}
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		// 使用滚动数组
		x, y := i&1, (i-1)&1
		// 注意还要跟 nums[i] 比较，因为 nums[i] 可能大于前几个的乘积
		// 如果没有负数，dp[i][min] 一直是 0
		dp[x][0] = max(dp[y][0]*nums[i], dp[y][1]*nums[i], nums[i])
		dp[x][1] = min(dp[y][0]*nums[i], dp[y][1]*nums[i], nums[i])
		res = max(res, dp[x][0])
	}
	return res
}

func max(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func min(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}
