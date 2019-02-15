package array

/*
	solution 1: sequential
	solution 2: binary
*/

/*
func findPeakElement(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return i - 1
		}
	}
	return len(nums) - 1
}
*/

func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
