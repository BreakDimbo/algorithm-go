package array

/*
	solution 1: hash map
	solution 2: 2∗(a+b+c)−(a+a+b+b+c)=c
	solution 3:
	If we take XOR of zero and some bit, it will return that bit
	a^0 =a
	If we take XOR of two same bits, it will return 0
	a^a = 0
	a^b^a = a^a^b = b
*/

func singleNumber(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}
	return a
}
