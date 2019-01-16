package recursive

/*
	solution 1: 排序后找第 n/2 个元素
	solution 2: 使用 map 找到出现次数最多的元素
	solution 3: Boyer-Moore Voting Algorithm - 因为多数元素的个数大于 n/2
*/

/* solution 1
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
*/

func majorityElement(nums []int) int {
	count, candidate := 0, nums[0]
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
