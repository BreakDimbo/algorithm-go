package bit

/*
	solution 1: x = x&(x-1)
	solution 2: >> and & 1
*/

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num &= (num - 1)
		count++
	}
	return count
}
