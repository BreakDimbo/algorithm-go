package string

import "math"

/*
	solution 1: 处理 edge，base = base * 10 + s[i]
*/

func myAtoi(str string) int {
	sign, base, i := 1, 0, 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i < len(str) && str[i] == '+' {
		i++
	} else if i < len(str) && str[i] == '-' {
		sign = -1
		i++
	}

	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		// handle overflow
		if base > math.MaxInt32/10 || base == math.MaxInt32/10 && int(str[i]-'0') > 7 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		base = base*10 + int(str[i]-'0')
		i++
	}
	return base * sign
}
