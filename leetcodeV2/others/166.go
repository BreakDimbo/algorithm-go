package others

import "fmt"

/*
	solution 1:
*/

func fractionToDecimal(numerator int, denominator int) string {

	var res string
	// handle negative number
	if (numerator>>31 ^ denominator>>31) == 1 {
		res += "-"
	}

	// in case of overflow of -2^32 / -1
	var num, den int64
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	num = int64(numerator)
	den = int64(denominator)

	// get the integer part
	integer := num / den
	res += fmt.Sprintf("%d", integer)
	rmd := num % den
	if rmd == 0 {
		return res
	}

	res += "."
	rmd *= 10
	rmdMap := make(map[int64]int)
	for rmd != 0 {
		quotation := rmd / den
		if index, ok := rmdMap[rmd]; ok {
			res = res[0:index] + "(" + res[index:]
			res += ")"
			break
		}

		rmdMap[rmd] = len(res)
		res += fmt.Sprintf("%d", quotation)
		rmd = (rmd % den) * 10
	}
	return res
}
