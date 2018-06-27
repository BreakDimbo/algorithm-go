package sort

import (
	"strings"
)

type Comparable interface {
	CompareTo(Comparable) int
}

type CString string

// if sc < c then return 0
func (cs CString) CompareTo(c Comparable) int {
	cStr, ok := c.(CString)
	if !ok {
		panic("wrong type of c, not ComparableString")
	}
	stredCs := string(cs)
	stredC := string(cStr)
	return strings.Compare(stredCs, stredC)
}
