package array

// sliding window not for negative

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	count, curSum := 0, 0
	for i < j && i < len(nums) {
		if curSum < k && j < len(nums) {
			curSum += nums[j]
			j++
		} else if curSum > k {
			curSum -= nums[i]
			i++
		}
		if curSum == k && i < len(nums) {
			count++
			curSum -= nums[i]
			i++
		}
	}
	return count
}
