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

// 二分法
func lengthOfLIs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		index := lowerBound(res, nums[i]) // 找到最小的比目标值大的数的下标
		if index > len(res)-1 {
			res = append(res, nums[i])
		} else {
			res[index] = nums[i] // 替换
		}
	}
	return len(res)
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
