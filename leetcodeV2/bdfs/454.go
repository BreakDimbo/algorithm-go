package bdfs

/*
	solution 1: brute force O(n^4)
	solution 2: hash (O(n^2))
*/

func fourSumCount(A []int, B []int, C []int, D []int) int {
	count := 0
	mp := make(map[int]int)
	for _, m := range C {
		for _, n := range D {
			mp[m+n]++
		}
	}
	for _, i := range A {
		for _, j := range B {
			if v, ok := mp[-(i + j)]; ok {
				count += v
			}
		}
	}
	return count
}
