package array

/*
	solution 1: sorting
	solution 2: hash
	solution 3: cycle detect
*/

func findDuplicate(nums []int) int {
	fast, slow := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			head := nums[0]
			for head != fast {
				head = nums[head]
				fast = nums[fast]
			}
			return fast
		}
	}
}
