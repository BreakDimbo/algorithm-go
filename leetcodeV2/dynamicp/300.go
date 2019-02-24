package dynamicp

/*
	solution 1: dp
		dp[i] contain i LIS

	solution 2: binary
*/

/*
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxGlobal := 0
	for i := 1; i < len(nums); i++ {
		maxValue := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxValue = util.Max(maxValue, dp[j])
			}
		}
		dp[i] = maxValue + 1
		maxGlobal = util.Max(maxGlobal, dp[i])
	}
	return maxGlobal
}
*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lis := make([]int, 0, len(nums))

	for _, num := range nums {
		if len(lis) == 0 || lis[len(lis)-1] < num {
			lis = append(lis, num)
		} else {
			i := bsearchFisrtLarge(lis, num)
			lis[i] = num
		}
	}
	return len(lis)
}

func bsearchFisrtLarge(lis []int, t int) int {
	lo, hi := 0, len(lis)-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)

		if lis[mid] < t {
			lo = mid + 1
		} else {
			if mid == 0 || lis[mid-1] < t {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
