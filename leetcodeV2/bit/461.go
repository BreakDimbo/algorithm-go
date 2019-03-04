package bit

/*
	solution 1: xor and find number of 1
*/

func hammingDistance(x int, y int) int {
	k := x ^ y
	count := 0
	for k != 0 {
		k &= k - 1
		count++
	}
	return count
}
