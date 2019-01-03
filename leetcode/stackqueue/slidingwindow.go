package stackqueue

import "container/list"

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	// window 存储的是元素在 nums 的下标值
	window := list.New()
	res := []int{}
	for i, v := range nums {
		// 此时元素未被加入，所以是 i >= k
		if i >= k && window.Front().Value.(int) <= i-k {
			window.Remove(window.Front())
		}
		for window.Len() != 0 && v >= nums[window.Back().Value.(int)] {
			window.Remove(window.Back())
		}
		window.PushBack(i)

		// 此时元素已经被加入，所以是 i >= k - 1
		if i >= k-1 {
			res = append(res, nums[window.Front().Value.(int)])
		}
	}
	return res
}
