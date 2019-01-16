package binary

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1:
	ensure:
	1) len(left_part) == len(right_part)
		i + j = m-i+n-j+1 => j = (m+n+1)/2-i
	2) max(left_part) <= min(right_part)
	3) i==0,j==0,i==m,j==n 的处理
	4) i < m ==> j > 0 and i > 0 ==> j < n
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2
	m, n := len(a), len(b)
	// need to promise n >= m
	if m > n {
		m, n, a, b = n, m, b, a
	}
	if n == 0 {
		return -1
	}

	// start the binay search for perfect i from [imin, imax]
	// a[0],a[1]..a[i-1] | a[i],a[i+1]...a[m-1]
	// b[0],b[1]..b[j-1] | b[j],b[j+1]...b[n]
	imin, imax, halfLength := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLength - i
		if i > 0 && a[i-1] > b[j] {
			imax = i - 1
		} else if i < m && a[i] < b[j-1] {
			imin = i + 1
		} else {
			// find the perfect i
			maxOfLeft, minOfRight := 0, 0
			if i == 0 {
				maxOfLeft = b[j-1]
			} else if j == 0 {
				maxOfLeft = a[i-1]
			} else {
				maxOfLeft = util.Max(a[i-1], b[j-1])
			}

			// if m+n is odd then max of left is result
			if (m+n)&1 == 1 {
				return float64(maxOfLeft)
			}

			if i == m {
				minOfRight = b[j]
			} else if j == n {
				minOfRight = a[i]
			} else {
				minOfRight = util.Min(a[i], b[j])
			}
			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}
	return -1
}
