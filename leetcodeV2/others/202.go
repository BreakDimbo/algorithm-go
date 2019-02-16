package others

/*
	solution 1: Floyd Cycle detection
	solution 2: hashMap
*/

func isHappy(n int) bool {
	memo := make(map[int]bool)

	for n != 1 && !memo[n] {
		memo[n] = true
		tmp := n
		for n != 0 {
			tmp += (n % 10) * (n % 10)
			n /= 10
		}
		n = tmp
	}

	if n == 1 {
		return true
	} else {
		return false
	}
}
