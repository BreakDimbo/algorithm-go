package others

/*
	solution 1: brute force
	solution 2: dp
		dp[i] 前 i 个节点的最大容量
		dp[i] = dp[i-1] + max([0,i]*h)

	solution 3: two point
		因为高度是由最短边决定的，每次移动较短的 point，保存最大面积
		当指针在两端的时候，只有移动较短边的指针才有可能获得更大的面积
		反证法
		We starts with the widest container, l = 0 and r = n - 1. Let's say the left one is shorter: h[l] < h[r].
		Then, this is already the largest container the left one can possibly form. There's no need to consider it again.
		Therefore, we just throw it away and start again with l = 1 and r = n -1.
*/
/* solution 1
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	max := 0

	for i := range height {
		for j := i + 1; j < len(height); j++ {
			h := util.Min(height[i], height[j])
			l := j - i
			if h*l > max {
				max = h * l
			}
		}
	}
	return max
}
*/

/* solution 2
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	dp := make([]int, len(height)+1)
	dp[0] = 0

	for i := 1; i < len(height); i++ {
		max := 0
		for j := 0; j < i; j++ {
			h := i - j
			l := util.Min(height[i], height[j])
			if h*l > max {
				max = h * l
			}
		}
		dp[i] = util.Max(dp[i-1], max)
	}
	return dp[len(height)-1]
}
*/

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right, max := 0, len(height)-1, 0

	for left < right {
		var s int
		if height[left] < height[right] {
			s = (right - left) * height[left]
			left++
		} else {
			s = (right - left) * height[right]
			right--
		}
		if s > max {
			max = s
		}
	}
	return max
}
