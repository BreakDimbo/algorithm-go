package array

/*
	solution 1: hash
	solution 2: mark number as negative
*/

func findDisappearedNumbers(nums []int) []int {
	h := make(map[int]bool)
	var res []int
	for _, num := range nums {
		h[num] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := h[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
