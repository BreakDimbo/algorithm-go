package array

/*
	solution 1: use stack
	solution 2: cut first n-k elements to the tail
	solution 3: reverse
*/

/*
func rotate(nums []int, k int) {
	k %= len(nums)
	tmp := nums
	for i := 0; i < len(nums)-k; i++ {
		e := tmp[0]
		tmp = append(tmp[1:], e)
	}

	for i := range tmp {
		nums[i] = tmp[i]
	}
}
*/

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
