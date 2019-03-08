package priorityQueue

func maxSlidingWindow(nums []int, k int) []int {
	// solution 1 : MaxHeap
	// solution 2: dequeue

	// 保存数组下标
	window := make([]int, 0, k)
	res := make([]int, 0)

	for i, num := range nums {
		// 判断左边是否需要移动
		if i >= k && i-k > window[0] {
			window = window[1:]
		}

		// 新加入元素从右向左依次与队列中元素比较，保证最左边元素是当前最大值
		for len(window) > 0 && nums[window[len(window)-1]] < num {
			window = window[:len(window)-1]
		}
		window = append(window, i)

		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
