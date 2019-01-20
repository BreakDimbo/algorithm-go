package array

/*
	solution 1: 检查相邻元素是否相同
	solution 2: 快慢指针
*/

/*
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	i := 1
	for i < n {
		if nums[i] == nums[i-1] {
			// move
			for k := i; k < n-1; k++ {
				nums[k] = nums[k+1]
			}
			n--
		} else {
			i++
		}
	}
	return n
}
*/

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 1 {
		return n
	}
	i, j := 0, 0
	for j != n-1 {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i
}
