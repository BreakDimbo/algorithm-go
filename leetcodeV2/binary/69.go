package binary

/*
	solution 1: binary
	solution 2: 牛顿迭代法
*/

/*
func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}

	l, r, res := 0, x, 0
	for l <= r {
		mid := l + ((r - l) >> 1)
		if mid == x/mid {
			return mid
		}
		if mid < x/mid {
			l = mid + 1
			res = mid
		} else {
			r = mid - 1
		}
	}
	return res
}
*/

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
