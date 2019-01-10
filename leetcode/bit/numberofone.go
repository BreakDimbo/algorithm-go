package bit

func hammingWeightSimple(num uint32) int {
	count := 0
	for num != 0 {
		if num%2 != 0 {
			count++
		}
		num = num >> 1
	}
	return count
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
