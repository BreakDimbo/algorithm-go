package array

/*
	solution 1: use swap keep constant space and also make use of the length of the array
		for any k positive numbers (duplicates allowed),
		the first missing positive number must be within [1,k+1]
*/

func firstMissingPositive(nums []int) int {
	for i := range nums {
		// 注意这里使用 for
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	i := 0
	for ; i < len(nums) && nums[i] == i+1; i++ {
	}
	return i + 1
}
