package binary

import (
	"math"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right, res := 0, x, 0
	for left <= right {
		mid := (left + right) / 2
		if mid == x/mid { // 小心 m * m 越界
			return mid
		}
		if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
			res = mid
		}
	}
	return res
}

func mySqrtN(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func sqrtF(x float64, n int) float64 {
	if x == 0 || x == 1 {
		return x
	}
	var mid, left float64
	right := x
	fac := math.Pow10(-n)

	for left <= right {
		mid := (left + right) / 2
		if math.Abs(mid-x/mid) <= fac { // 注意这里使用绝对值
			return mid
		}
		if mid < x/mid {
			left = mid
		} else {
			right = mid
		}
	}
	return mid
}
