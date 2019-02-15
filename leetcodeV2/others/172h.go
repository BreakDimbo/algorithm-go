package others

/*
	solution 1: divide by 5
		n 个数中有多少个 5 的倍数
			n/5 最多是 5 的几倍，就有几个 5 的倍数
			处理 5^x 额外多一个 5
*/

func trailingZeroes(n int) int {
	res := 0
	for n != 0 {
		tmp := n / 5
		res += tmp
		n = tmp
	}
	return res
}
