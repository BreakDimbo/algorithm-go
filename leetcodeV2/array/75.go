package array

/*
	solution 1: counting 0,1,2
	solution 2: sort
	solution 3: swap
*/

/*
func sortColors(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
*/

func sortColors(nums []int) {
	j, k := 0, len(nums)-1
	for i := 0; i <= k; i++ {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		if nums[i] == 2 {
			nums[i], nums[k] = nums[k], nums[i]
			i--
			k--
		}
	}
}
