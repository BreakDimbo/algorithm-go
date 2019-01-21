package others

import "math"

/*
	solution 1: 使用减法,注意正负号
	solution 2: bit
*/

/*
func divide(dividend int, divisor int) int {
	minus := false
	if dividend < 0 && divisor < 0 {
		if dividend <= math.MinInt32 && divisor == -1 {
			return math.MaxInt32
		}
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 {
		dividend = -dividend
		minus = true
	} else if divisor < 0 {
		divisor = -divisor
		minus = true
	}
	k := 0

	for dividend >= 0 {
		dividend -= divisor
		k++
	}
	if minus {
		return -(k - 1)
	}
	return k - 1
}
*/

func divide(dividend int, divisor int) int {
	if dividend <= math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		sign = -1
	}
	did := int(math.Abs(float64(dividend)))
	div := int(math.Abs(float64(divisor)))
	res := 0
	for did >= div {
		tmp, multiple := div, 1
		for did >= tmp<<1 {
			tmp <<= 1
			multiple <<= 1
		}
		did -= tmp
		res += multiple
	}
	return res * sign
}
