package array

/*
	solution 1: two pointers
*/

func moveZeroes(nums []int) {
	i, j := 0, 0
	for {
		for i < len(nums) && nums[i] != 0 {
			i++
		}
		for j < len(nums) && (nums[j] == 0 || j <= i) {
			j++
		}
		if j == len(nums) {
			break
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
