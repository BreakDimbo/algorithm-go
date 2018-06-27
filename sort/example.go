package sort

import (
	"fmt"
)

type Example struct {
}

func (e *Example) Sort(c []Comparable) {
	// implement by the specific sort method
}

func (e *Example) IsSorted(c []Comparable) bool {
	for i := 0; i < len(c); i++ {
		if e.less(c[i+1], c[i]) {
			return false
		}
	}
	return true
}

func (e *Example) show(c []Comparable) {
	for i := 0; i < len(c); i++ {
		fmt.Printf("%s, ", c[i])
	}
	fmt.Println()
}

func (e *Example) less(v Comparable, w Comparable) bool {
	return v.CompareTo(w) < 0
}

func (e *Example) exch(c []Comparable, i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}
