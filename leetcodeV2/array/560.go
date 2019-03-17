package array

// sliding window not for negative
// what we need is [0, i - 1] and SUM[0, j] to get SUM[i, j] == k

func subarraySum(nums []int, k int) int {
	var sum, count int
	h := make(map[int]int)
	h[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := h[sum-k]; ok {
			count += v
		}
		h[sum]++
	}
	return count
}
