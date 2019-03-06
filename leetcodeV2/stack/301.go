package stack

import "math"

/*
	solution 1: brute force
	solution 2: limit backtracing
*/

var (
	res         map[string]bool
	minRmvCount int
)

func removeInvalidParentheses(s string) []string {
	res = make(map[string]bool)
	minRmvCount = math.MaxInt32
	backtraceRmv(s, 0, 0, 0, 0, "")
	results := make([]string, 0, len(res))
	for r := range res {
		results = append(results, r)
	}
	return results
}

func backtraceRmv(s string, index, lc, rc, rmc int, expr string) {
	// find the valid result
	if index == len(s) {
		if lc == rc && rmc <= minRmvCount {
			if rmc < minRmvCount {
				res = make(map[string]bool)
				minRmvCount = rmc
			}
			res[expr] = true
		}
		return
	}

	// if not ( & ) just put into the res and recursive
	if s[index] != '(' && s[index] != ')' {
		backtraceRmv(s, index+1, lc, rc, rmc, expr+string(s[index]))
	} else {
		// delete s[index]
		backtraceRmv(s, index+1, lc, rc, rmc+1, expr)

		// put into expr
		if s[index] == '(' {
			backtraceRmv(s, index+1, lc+1, rc, rmc, expr+string(s[index]))
		} else if lc > rc {
			backtraceRmv(s, index+1, lc, rc+1, rmc, expr+string(s[index]))
		}
	}
}
