package array

/*
	solution 1: add 1 to the last ele and check if exceed 10
*/

func plusOne(digits []int) []int {
	i := len(digits) - 1
	digits[i] += 1
	for digits[i] > 9 {
		digits[i] = 0
		i--

		if i < 0 {
			digits = make([]int, len(digits)+1)
			digits[0] = 1
			break
		}
		digits[i] += 1
	}
	return digits
}
